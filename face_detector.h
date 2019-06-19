#pragma once

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
	IMAGE_LOAD_ERROR,
	SERIALIZATION_ERROR,
	UNKNOWN_ERROR,
} err_code;

typedef struct facedetector {
	void* cls;
	const char* err_str;
	err_code err_code;
} facedetector;

facedetector* facedetector_init();
void facedetector_free(facedetector* rec);

typedef struct returnvalue {
	long* rectangles;
	int length;
	const char* err_str;
	err_code err_code;
} returnvalue;

returnvalue* facedetector_detect(facedetector *det, const uint8_t* img_data, int len);

#ifdef __cplusplus
}
#endif
