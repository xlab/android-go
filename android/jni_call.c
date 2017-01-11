#include "jni_call.h"

jint JNI_DestroyJavaVM(JavaVM* vm) {
	return (*vm)->DestroyJavaVM(vm);
}

jint JNI_AttachCurrentThread(JavaVM* vm, JNIEnv** p_env, JavaVMAttachArgs* thr_args) {
	return (*vm)->AttachCurrentThread(vm, p_env, thr_args);
}

jint JNI_DetachCurrentThread(JavaVM* vm) {
	return (*vm)->DetachCurrentThread(vm);
}

jint JNI_GetEnv(JavaVM* vm, JNIEnv** p_env, jint version) {
	return (*vm)->GetEnv(vm, (void**)p_env, version);
}

jint JNI_AttachCurrentThreadAsDaemon(JavaVM* vm, JNIEnv** p_env, void* thr_args) {
	return (*vm)->AttachCurrentThreadAsDaemon(vm, p_env, thr_args);
}

jint JNIEnv_GetVersion(JNIEnv* env) {
	return (*env)->GetVersion(env);
}

jclass JNIEnv_DefineClass(JNIEnv* env, const char* name, jobject obj, jbyte* buf, jsize bufLen) {
	return (*env)->DefineClass(env, name, obj, buf, bufLen);
}

jclass JNIEnv_FindClass(JNIEnv* env, const char* name) {
	return (*env)->FindClass(env, name);
}

jmethodID JNIEnv_FromReflectedMethod(JNIEnv* env, jobject obj) {
	return (*env)->FromReflectedMethod(env, obj);
}

jfieldID JNIEnv_FromReflectedField(JNIEnv* env, jobject obj) {
	return (*env)->FromReflectedField(env, obj);
}

jobject JNIEnv_ToReflectedMethod(JNIEnv* env, jclass clazz, jmethodID id, jboolean isStatic) {
	return (*env)->ToReflectedMethod(env, clazz, id, isStatic);
}

jclass JNIEnv_GetSuperclass(JNIEnv* env, jclass clazz) {
	return (*env)->GetSuperclass(env, clazz);
}

jboolean JNIEnv_IsAssignableFrom(JNIEnv* env, jclass clazz1, jclass clazz2) {
	return (*env)->IsAssignableFrom(env, clazz1, clazz2);
}

jobject JNIEnv_ToReflectedField(JNIEnv* env, jclass clazz, jfieldID id, jboolean isStatic) {
	return (*env)->ToReflectedField(env, clazz, id, isStatic);
}

jint JNIEnv_Throw(JNIEnv* env, jthrowable ex) {
	return (*env)->Throw(env, ex);
}

jint JNIEnv_ThrowNew(JNIEnv* env, jclass clazz, const char* msg) {
	return (*env)->ThrowNew(env, clazz, msg);
}

jthrowable JNIEnv_ExceptionOccurred(JNIEnv* env) {
	return (*env)->ExceptionOccurred(env);
}

void JNIEnv_ExceptionDescribe(JNIEnv* env) {
	(*env)->ExceptionDescribe(env);
}

void JNIEnv_ExceptionClear(JNIEnv* env) {
	(*env)->ExceptionClear(env);
}

void JNIEnv_FatalError(JNIEnv* env, const char* msg) {
	(*env)->FatalError(env, msg);
}

jint JNIEnv_PushLocalFrame(JNIEnv* env, jint capacity) {
	return (*env)->PushLocalFrame(env, capacity);
}

jobject JNIEnv_PopLocalFrame(JNIEnv* env, jobject obj) {
	return (*env)->PopLocalFrame(env, obj);
}

jobject JNIEnv_NewGlobalRef(JNIEnv* env, jobject ref) {
	return (*env)->NewGlobalRef(env, ref);
}

void JNIEnv_DeleteGlobalRef(JNIEnv* env, jobject ref) {
	(*env)->DeleteGlobalRef(env, ref);
}

void JNIEnv_DeleteLocalRef(JNIEnv* env, jobject ref) {
	(*env)->DeleteLocalRef(env, ref);
}

jboolean JNIEnv_IsSameObject(JNIEnv* env, jobject ref1, jobject ref2) {
	return (*env)->IsSameObject(env, ref1, ref2);
}

jobject JNIEnv_NewLocalRef(JNIEnv* env, jobject obj) {
	return (*env)->NewLocalRef(env, obj);
}

jint JNIEnv_EnsureLocalCapacity(JNIEnv* env, jint capacity) {
	return (*env)->EnsureLocalCapacity(env, capacity);
}

jobject JNIEnv_AllocObject(JNIEnv* env, jclass clazz) {
	return (*env)->AllocObject(env, clazz);
}

jobject JNIEnv_NewObject(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->NewObjectA(env, clazz, id, args);
}

jclass JNIEnv_GetObjectClass(JNIEnv* env, jobject obj) {
	return (*env)->GetObjectClass(env, obj);
}

jboolean JNIEnv_IsInstanceOf(JNIEnv* env, jobject obj, jclass clazz) {
	return (*env)->IsInstanceOf(env, obj, clazz);
}

jmethodID JNIEnv_GetMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig) {
	return (*env)->GetMethodID(env, clazz, name, sig);
}

jobject JNIEnv_CallObjectMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallObjectMethodA(env,  obj, id, args);
}

jboolean JNIEnv_CallBooleanMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallBooleanMethodA(env, obj, id, args);
}

jbyte JNIEnv_CallByteMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallByteMethodA(env, obj, id, args);
}

jchar JNIEnv_CallCharMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallCharMethodA(env, obj, id, args);
}

jshort JNIEnv_CallShortMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallShortMethodA(env, obj, id, args);
}

jint JNIEnv_CallIntMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallIntMethodA(env, obj, id, args);
}

jlong JNIEnv_CallLongMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallLongMethodA(env, obj, id, args);
}

jfloat JNIEnv_CallFloatMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallFloatMethodA(env, obj, id, args);
}

jdouble JNIEnv_CallDoubleMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	return (*env)->CallDoubleMethodA(env, obj, id, args);
}

void JNIEnv_CallVoidMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args) {
	(*env)->CallVoidMethodA(env, obj, id, args);
}

jobject JNIEnv_CallNonvirtualObjectMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualObjectMethodA(env, obj, clazz, id, args);
}

jboolean JNIEnv_CallNonvirtualBooleanMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualBooleanMethodA(env, obj, clazz, id, args);
}

jbyte JNIEnv_CallNonvirtualByteMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualByteMethodA(env, obj, clazz, id, args);
}

jchar JNIEnv_CallNonvirtualCharMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualCharMethodA(env, obj, clazz, id, args);
}

jshort JNIEnv_CallNonvirtualShortMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualShortMethodA(env, obj, clazz, id, args);
}

jint JNIEnv_CallNonvirtualIntMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualIntMethodA(env, obj, clazz, id, args);
}

jlong JNIEnv_CallNonvirtualLongMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualLongMethodA(env, obj, clazz, id, args);
}

jfloat JNIEnv_CallNonvirtualFloatMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualFloatMethodA(env, obj, clazz, id, args);
}

jdouble JNIEnv_CallNonvirtualDoubleMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallNonvirtualDoubleMethodA(env, obj, clazz, id, args);
}

void JNIEnv_CallNonvirtualVoidMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args) {
	(*env)->CallNonvirtualVoidMethodA(env, obj, clazz, id, args);
}

jfieldID JNIEnv_GetFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig) {
	return (*env)->GetFieldID(env, clazz, name, sig);
}

jobject JNIEnv_GetObjectField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetObjectField(env, obj, id);
}

jboolean JNIEnv_GetBooleanField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetBooleanField(env, obj, id);
}

jbyte JNIEnv_GetByteField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetByteField(env, obj, id);
}

jchar JNIEnv_GetCharField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetCharField(env, obj, id);
}

jshort JNIEnv_GetShortField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetShortField(env, obj, id);
}

jint JNIEnv_GetIntField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetIntField(env, obj, id);
}

jlong JNIEnv_GetLongField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetLongField(env, obj, id);
}

jfloat JNIEnv_GetFloatField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetFloatField(env, obj, id);
}

jdouble JNIEnv_GetDoubleField(JNIEnv* env, jobject obj, jfieldID id) {
	return (*env)->GetDoubleField(env, obj, id);
}

void JNIEnv_SetObjectField(JNIEnv* env, jobject obj, jfieldID id, jobject val) {
	(*env)->SetObjectField(env, obj, id, val);
}

void JNIEnv_SetBooleanField(JNIEnv* env, jobject obj, jfieldID id, jboolean val) {
	(*env)->SetBooleanField(env, obj, id, val);
}

void JNIEnv_SetByteField(JNIEnv* env, jobject obj, jfieldID id, jbyte val) {
	(*env)->SetByteField(env, obj, id, val);
}

void JNIEnv_SetCharField(JNIEnv* env, jobject obj, jfieldID id, jchar val) {
	(*env)->SetCharField(env, obj, id, val);
}

void JNIEnv_SetShortField(JNIEnv* env, jobject obj, jfieldID id, jshort val) {
	(*env)->SetShortField(env, obj, id, val);
}

void JNIEnv_SetIntField(JNIEnv* env, jobject obj, jfieldID id, jint val) {
	(*env)->SetIntField(env, obj, id, val);
}

void JNIEnv_SetLongField(JNIEnv* env, jobject obj, jfieldID id, jlong val) {
	(*env)->SetLongField(env, obj, id, val);
}

void JNIEnv_SetFloatField(JNIEnv* env, jobject obj, jfieldID id, jfloat val) {
	(*env)->SetFloatField(env, obj, id, val);
}

void JNIEnv_SetDoubleField(JNIEnv* env, jobject obj, jfieldID id, jdouble val) {
	(*env)->SetDoubleField(env, obj, id, val);
}

jmethodID JNIEnv_GetStaticMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig) {
	return (*env)->GetStaticMethodID(env, clazz, name, sig);
}

jobject JNIEnv_CallStaticObjectMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticObjectMethodA(env, clazz, id, args);
}

jboolean JNIEnv_CallStaticBooleanMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticBooleanMethodA(env, clazz, id, args);
}

jbyte JNIEnv_CallStaticByteMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticByteMethodA(env, clazz, id, args);
}

jchar JNIEnv_CallStaticCharMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticCharMethodA(env, clazz, id, args);
}

jshort JNIEnv_CallStaticShortMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticShortMethodA(env, clazz, id, args);
}

jint JNIEnv_CallStaticIntMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticIntMethodA(env, clazz, id, args);
}

jlong JNIEnv_CallStaticLongMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticLongMethodA(env, clazz, id, args);
}

jfloat JNIEnv_CallStaticFloatMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticFloatMethodA(env, clazz, id, args);
}

jdouble JNIEnv_CallStaticDoubleMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	return (*env)->CallStaticDoubleMethodA(env, clazz, id, args);
}

void JNIEnv_CallStaticVoidMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args) {
	(*env)->CallStaticVoidMethodA(env, clazz, id, args);
}

jfieldID JNIEnv_GetStaticFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig) {
	return (*env)->GetStaticFieldID(env, clazz, name, sig);
}

jobject JNIEnv_GetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticObjectField(env, clazz, id);
}

jboolean JNIEnv_GetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticBooleanField(env, clazz, id);
}

jbyte JNIEnv_GetStaticByteField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticByteField(env, clazz, id);
}

jchar JNIEnv_GetStaticCharField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticCharField(env, clazz, id);
}

jshort JNIEnv_GetStaticShortField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticShortField(env, clazz, id);
}

jint JNIEnv_GetStaticIntField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticIntField(env, clazz, id);
}

jlong JNIEnv_GetStaticLongField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticLongField(env, clazz, id);
}

jfloat JNIEnv_GetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticFloatField(env, clazz, id);
}

jdouble JNIEnv_GetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID id) {
	return (*env)->GetStaticDoubleField(env, clazz, id);
}

void JNIEnv_SetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID id, jobject val) {
	(*env)->SetStaticObjectField(env, clazz, id, val);
}

void JNIEnv_SetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID id, jboolean val) {
	(*env)->SetStaticBooleanField(env, clazz, id, val);
}

void JNIEnv_SetStaticByteField(JNIEnv* env, jclass clazz, jfieldID id, jbyte val) {
	(*env)->SetStaticByteField(env, clazz, id, val);
}

void JNIEnv_SetStaticCharField(JNIEnv* env, jclass clazz, jfieldID id, jchar val) {
	(*env)->SetStaticCharField(env, clazz, id, val);
}

void JNIEnv_SetStaticShortField(JNIEnv* env, jclass clazz, jfieldID id, jshort val) {
	(*env)->SetStaticShortField(env, clazz, id, val);
}

void JNIEnv_SetStaticIntField(JNIEnv* env, jclass clazz, jfieldID id, jint val) {
	(*env)->SetStaticIntField(env, clazz, id, val);
}

void JNIEnv_SetStaticLongField(JNIEnv* env, jclass clazz, jfieldID id, jlong val) {
	(*env)->SetStaticLongField(env, clazz, id, val);
}

void JNIEnv_SetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID id, jfloat val) {
	(*env)->SetStaticFloatField(env, clazz, id, val);
}

void JNIEnv_SetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID id, jdouble val) {
	(*env)->SetStaticDoubleField(env, clazz, id, val);
}

jstring JNIEnv_NewString(JNIEnv* env, jchar* buf, jsize bufLen) {
	return (*env)->NewString(env, buf, bufLen);
}

jsize JNIEnv_GetStringLength(JNIEnv* env, jstring str) {
	return (*env)->GetStringLength(env, str);
}

const jchar* JNIEnv_GetStringChars(JNIEnv* env, jstring str, jboolean* isCopy) {
	return (*env)->GetStringChars(env, str, isCopy);
}

void JNIEnv_ReleaseStringChars(JNIEnv* env, jstring str, jchar* chars) {
	(*env)->ReleaseStringChars(env, str, chars);
}

jstring JNIEnv_NewStringUTF(JNIEnv* env, const char* str) {
	return (*env)->NewStringUTF(env, str);
}

jsize JNIEnv_GetStringUTFLength(JNIEnv* env, jstring str) {
	return (*env)->GetStringUTFLength(env, str);
}

const char* JNIEnv_GetStringUTFChars(JNIEnv* env, jstring str, jboolean* isCopy) {
	return (*env)->GetStringUTFChars(env, str, isCopy);
}

void JNIEnv_ReleaseStringUTFChars(JNIEnv* env, jstring str, const char* utf) {
	(*env)->ReleaseStringUTFChars(env, str, utf);
}

jsize JNIEnv_GetArrayLength(JNIEnv* env, jarray arr) {
	return (*env)->GetArrayLength(env, arr);
}

jobjectArray JNIEnv_NewObjectArray(JNIEnv* env, jsize length, jclass clazz, jobject obj) {
	return (*env)->NewObjectArray(env, length, clazz, obj);
}

jobject JNIEnv_GetObjectArrayElement(JNIEnv* env, jobjectArray arr, jsize index) {
	return (*env)->GetObjectArrayElement(env, arr, index);
}

void JNIEnv_SetObjectArrayElement(JNIEnv* env, jobjectArray arr, jsize index, jobject obj) {
	(*env)->SetObjectArrayElement(env, arr, index, obj);
}

jbooleanArray JNIEnv_NewBooleanArray(JNIEnv* env, jsize length) {
	return (*env)->NewBooleanArray(env, length);
}

jbyteArray JNIEnv_NewByteArray(JNIEnv* env, jsize length) {
	return (*env)->NewByteArray(env, length);
}

jcharArray JNIEnv_NewCharArray(JNIEnv* env, jsize length) {
	return (*env)->NewCharArray(env, length);
}

jshortArray JNIEnv_NewShortArray(JNIEnv* env, jsize length) {
	return (*env)->NewShortArray(env, length);
}

jintArray JNIEnv_NewIntArray(JNIEnv* env, jsize length) {
	return (*env)->NewIntArray(env, length);
}

jlongArray JNIEnv_NewLongArray(JNIEnv* env, jsize length) {
	return (*env)->NewLongArray(env, length);
}

jfloatArray JNIEnv_NewFloatArray(JNIEnv* env, jsize length) {
	return (*env)->NewFloatArray(env, length);
}

jdoubleArray JNIEnv_NewDoubleArray(JNIEnv* env, jsize length) {
	return (*env)->NewDoubleArray(env, length);
}

jboolean* JNIEnv_GetBooleanArrayElements(JNIEnv* env, jbooleanArray arr, jboolean* isCopy) {
	return (*env)->GetBooleanArrayElements(env, arr, isCopy);
}

jbyte* JNIEnv_GetByteArrayElements(JNIEnv* env, jbyteArray arr, jboolean* isCopy) {
	return (*env)->GetByteArrayElements(env, arr, isCopy);
}

jchar* JNIEnv_GetCharArrayElements(JNIEnv* env, jcharArray arr, jboolean* isCopy) {
	return (*env)->GetCharArrayElements(env, arr, isCopy);
}

jshort* JNIEnv_GetShortArrayElements(JNIEnv* env, jshortArray arr, jboolean* isCopy) {
	return (*env)->GetShortArrayElements(env, arr, isCopy);
}

jint* JNIEnv_GetIntArrayElements(JNIEnv* env, jintArray arr, jboolean* isCopy) {
	return (*env)->GetIntArrayElements(env, arr, isCopy);
}

jlong* JNIEnv_GetLongArrayElements(JNIEnv* env, jlongArray arr, jboolean* isCopy) {
	return (*env)->GetLongArrayElements(env, arr, isCopy);
}

jfloat* JNIEnv_GetFloatArrayElements(JNIEnv* env, jfloatArray arr, jboolean* isCopy) {
	return (*env)->GetFloatArrayElements(env, arr, isCopy);
}

jdouble* JNIEnv_GetDoubleArrayElements(JNIEnv* env, jdoubleArray arr, jboolean* isCopy) {
	return (*env)->GetDoubleArrayElements(env, arr, isCopy);
}

void JNIEnv_ReleaseBooleanArrayElements(JNIEnv* env, jbooleanArray arr, jboolean* elems, jint mode) {
	(*env)->ReleaseBooleanArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseByteArrayElements(JNIEnv* env, jbyteArray arr, jbyte* elems, jint mode) {
	(*env)->ReleaseByteArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseCharArrayElements(JNIEnv* env, jcharArray arr, jchar* elems, jint mode) {
	(*env)->ReleaseCharArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseShortArrayElements(JNIEnv* env, jshortArray arr, jshort* elems, jint mode) {
	(*env)->ReleaseShortArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseIntArrayElements(JNIEnv* env, jintArray arr, jint* elems, jint mode) {
	(*env)->ReleaseIntArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseLongArrayElements(JNIEnv* env, jlongArray arr, jlong* elems, jint mode) {
	(*env)->ReleaseLongArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseFloatArrayElements(JNIEnv* env, jfloatArray arr, jfloat* elems, jint mode) {
	(*env)->ReleaseFloatArrayElements(env, arr, elems, mode);
}

void JNIEnv_ReleaseDoubleArrayElements(JNIEnv* env, jdoubleArray arr, jdouble* elems, jint mode) {
	(*env)->ReleaseDoubleArrayElements(env, arr, elems, mode);
}

void JNIEnv_GetBooleanArrayRegion(JNIEnv* env, jbooleanArray arr, jsize start, jsize length, jboolean* buf) {
	(*env)->GetBooleanArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetByteArrayRegion(JNIEnv* env, jbyteArray arr, jsize start, jsize length, jbyte* buf) {
	(*env)->GetByteArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetCharArrayRegion(JNIEnv* env, jcharArray arr, jsize start, jsize length, jchar* buf) {
	(*env)->GetCharArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetShortArrayRegion(JNIEnv* env, jshortArray arr, jsize start, jsize length, jshort* buf) {
	(*env)->GetShortArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetIntArrayRegion(JNIEnv* env, jintArray arr, jsize start, jsize length, jint* buf) {
	(*env)->GetIntArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetLongArrayRegion(JNIEnv* env, jlongArray arr, jsize start, jsize length, jlong* buf) {
	(*env)->GetLongArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetFloatArrayRegion(JNIEnv* env, jfloatArray arr, jsize start, jsize length, jfloat* buf) {
	(*env)->GetFloatArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_GetDoubleArrayRegion(JNIEnv* env, jdoubleArray arr, jsize start, jsize length, jdouble* buf) {
	(*env)->GetDoubleArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetBooleanArrayRegion(JNIEnv* env, jbooleanArray arr, jsize start, jsize length, jboolean* buf) {
	(*env)->SetBooleanArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetByteArrayRegion(JNIEnv* env, jbyteArray arr, jsize start, jsize length, jbyte* buf) {
	(*env)->SetByteArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetCharArrayRegion(JNIEnv* env, jcharArray arr, jsize start, jsize length, jchar* buf) {
	(*env)->SetCharArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetShortArrayRegion(JNIEnv* env, jshortArray arr, jsize start, jsize length, jshort* buf) {
	(*env)->SetShortArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetIntArrayRegion(JNIEnv* env, jintArray arr, jsize start, jsize length, jint* buf) {
	(*env)->SetIntArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetLongArrayRegion(JNIEnv* env, jlongArray arr, jsize start, jsize length, jlong* buf) {
	(*env)->SetLongArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetFloatArrayRegion(JNIEnv* env, jfloatArray arr, jsize start, jsize length, jfloat* buf) {
	(*env)->SetFloatArrayRegion(env, arr, start, length, buf);
}

void JNIEnv_SetDoubleArrayRegion(JNIEnv* env, jdoubleArray arr, jsize start, jsize length, jdouble* buf) {
	(*env)->SetDoubleArrayRegion(env, arr, start, length, buf);
}

jint JNIEnv_RegisterNatives(JNIEnv* env, jclass clazz, JNINativeMethod* methods, jint nMethods) {
	return (*env)->RegisterNatives(env, clazz, methods, nMethods);
}

jint JNIEnv_UnregisterNatives(JNIEnv* env, jclass clazz) {
	return (*env)->UnregisterNatives(env, clazz);
}

jint JNIEnv_MonitorEnter(JNIEnv* env, jobject obj) {
	return (*env)->MonitorEnter(env, obj);
}

jint JNIEnv_MonitorExit(JNIEnv* env, jobject obj) {
	return (*env)->MonitorExit(env, obj);
}

jint JNIEnv_GetJavaVM(JNIEnv* env, JavaVM** p_vm) {
	return (*env)->GetJavaVM(env, p_vm);
}

void JNIEnv_GetStringRegion(JNIEnv* env, jstring str, jsize start, jsize length, jchar* buf) {
	(*env)->GetStringRegion(env, str, start, length, buf);
}

void JNIEnv_GetStringUTFRegion(JNIEnv* env, jstring str, jsize start, jsize length, char* buf) {
	(*env)->GetStringUTFRegion(env, str, start, length, buf);
}

void* JNIEnv_GetPrimitiveArrayCritical(JNIEnv* env, jarray arr, jboolean* isCopy) {
	return (*env)->GetPrimitiveArrayCritical(env, arr, isCopy);
}

void JNIEnv_ReleasePrimitiveArrayCritical(JNIEnv* env, jarray arr, void* carray, jint mode) {
	(*env)->ReleasePrimitiveArrayCritical(env, arr, carray, mode);
}

const jchar* JNIEnv_GetStringCritical(JNIEnv* env, jstring str, jboolean* isCopy) {
	return (*env)->GetStringCritical(env, str, isCopy);
}

void JNIEnv_ReleaseStringCritical(JNIEnv* env, jstring str, jchar* carray) {
	(*env)->ReleaseStringCritical(env, str, carray);
}

jweak JNIEnv_NewWeakGlobalRef(JNIEnv* env, jobject obj) {
	return (*env)->NewWeakGlobalRef(env, obj);
}

void JNIEnv_DeleteWeakGlobalRef(JNIEnv* env, jweak obj) {
	(*env)->DeleteWeakGlobalRef(env, obj);
}

jboolean JNIEnv_ExceptionCheck(JNIEnv* env) {
	return (*env)->ExceptionCheck(env);
}

jobject JNIEnv_NewDirectByteBuffer(JNIEnv* env, void* buf, jlong capacity) {
	return (*env)->NewDirectByteBuffer(env, buf, capacity);
}

void* JNIEnv_GetDirectBufferAddress(JNIEnv* env, jobject buf) {
	return (*env)->GetDirectBufferAddress(env, buf);
}

jlong JNIEnv_GetDirectBufferCapacity(JNIEnv* env, jobject buf) {
	return (*env)->GetDirectBufferCapacity(env, buf);
}

jobjectRefType JNIEnv_GetObjectRefType(JNIEnv* env, jobject obj) {
	return (*env)->GetObjectRefType(env, obj);
}

