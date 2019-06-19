#include <shared_mutex>
#include <dlib/image_loader/image_loader.h>
#include <dlib/image_processing/frontal_face_detector.h>
#include "face_detector.h"
#include "jpeg_mem_loader.h"

using namespace dlib;

static const size_t RECT_LEN = 4;
static const size_t RECT_SIZE = RECT_LEN * sizeof(long);

class FaceDetector
{
public:
	FaceDetector()
	{
		detector_ = get_frontal_face_detector();
	}

	std::vector<rectangle> detect(const matrix<rgb_pixel> &img)
	{
		std::lock_guard<std::mutex> lock(detector_mutex_);
		std::vector<rectangle> rects = detector_(img);
		return std::move(rects);
	}

private:
	std::mutex detector_mutex_;
	frontal_face_detector detector_;
};

// Plain C interface for Go.
facedetector *facedetector_init()
{
	facedetector *rec = (facedetector *)calloc(1, sizeof(facedetector));
	try
	{
		FaceDetector *cls = new FaceDetector();
		rec->cls = (void *)cls;
	}
	catch (serialization_error &e)
	{
		rec->err_str = strdup(e.what());
		rec->err_code = SERIALIZATION_ERROR;
	}
	catch (std::exception &e)
	{
		rec->err_str = strdup(e.what());
		rec->err_code = UNKNOWN_ERROR;
	}
	return rec;
}

returnvalue* facedetector_detect(facedetector *det, const uint8_t* img_data, int len)
{
	returnvalue* ret = (returnvalue*)calloc(1, sizeof(returnvalue));
	FaceDetector* cls = (FaceDetector*)(det->cls);
	matrix<rgb_pixel> img;
	load_mem_jpeg(img, img_data, len);

	std::vector<rectangle> rects = cls->detect(img);

	int size = rects.size();
	if (size == 0) {
		return ret;
	}
	ret->length = size;

	ret->rectangles = (long*)malloc(size * RECT_SIZE);
	for (int i = 0; i < size; i++) {
		long* dst = ret->rectangles + i * 4;
		dst[0] = rects[i].left();
		dst[1] = rects[i].top();
		dst[2] = rects[i].right();
		dst[3] = rects[i].bottom();
	}
	return ret;
}

void facedetector_free(facedetector *rec)
{
	if (rec->cls)
	{
		FaceDetector *cls = (FaceDetector *)(rec->cls);
		delete cls;
		rec->cls = NULL;
	}
	free(rec);
}
