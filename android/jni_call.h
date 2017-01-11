#ifndef JNI_CALL_H
#define JNI_CALL_H

#include <jni.h>

jint JNI_DestroyJavaVM(JavaVM* vm);
jint JNI_AttachCurrentThread(JavaVM* vm, JNIEnv** p_env, JavaVMAttachArgs* thr_args);
jint JNI_DetachCurrentThread(JavaVM* vm);
jint JNI_GetEnv(JavaVM* vm, JNIEnv** p_env, jint version);
jint JNI_AttachCurrentThreadAsDaemon(JavaVM* vm, JNIEnv** p_env, void* thr_args);

jint JNIEnv_GetVersion(JNIEnv* env);

jclass JNIEnv_DefineClass(JNIEnv* env, const char* name, jobject obj, jbyte* buf, jsize bufLen);
jclass JNIEnv_FindClass(JNIEnv* env, const char* name);

jmethodID JNIEnv_FromReflectedMethod(JNIEnv* env, jobject obj);
jfieldID JNIEnv_FromReflectedField(JNIEnv* env, jobject obj);
jobject JNIEnv_ToReflectedMethod(JNIEnv* env, jclass clazz, jmethodID id, jboolean isStatic);

jclass JNIEnv_GetSuperclass(JNIEnv* env, jclass clazz);
jboolean JNIEnv_IsAssignableFrom(JNIEnv* env, jclass clazz1, jclass clazz2);

jobject JNIEnv_ToReflectedField(JNIEnv* env, jclass clazz, jfieldID id, jboolean isStatic);

jint JNIEnv_Throw(JNIEnv* env, jthrowable ex);
jint JNIEnv_ThrowNew(JNIEnv* env, jclass clazz, const char* msg);
jthrowable JNIEnv_ExceptionOccurred(JNIEnv* env);
void JNIEnv_ExceptionDescribe(JNIEnv* env);
void JNIEnv_ExceptionClear(JNIEnv* env);
void JNIEnv_FatalError(JNIEnv* env, const char* msg);

jint JNIEnv_PushLocalFrame(JNIEnv* env, jint capacity);
jobject JNIEnv_PopLocalFrame(JNIEnv* env, jobject obj);

jobject JNIEnv_NewGlobalRef(JNIEnv* env, jobject ref);
void JNIEnv_DeleteGlobalRef(JNIEnv* env, jobject ref);
void JNIEnv_DeleteLocalRef(JNIEnv* env, jobject ref);
jboolean JNIEnv_IsSameObject(JNIEnv* env, jobject ref1, jobject ref2);

jobject JNIEnv_NewLocalRef(JNIEnv* env, jobject obj);
jint JNIEnv_EnsureLocalCapacity(JNIEnv* env, jint capacity);

jobject JNIEnv_AllocObject(JNIEnv* env, jclass clazz);
jobject JNIEnv_NewObject(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);

jclass JNIEnv_GetObjectClass(JNIEnv* env, jobject obj);
jboolean JNIEnv_IsInstanceOf(JNIEnv* env, jobject obj, jclass clazz);
jmethodID JNIEnv_GetMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

jobject JNIEnv_CallObjectMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jboolean JNIEnv_CallBooleanMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jbyte JNIEnv_CallByteMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jchar JNIEnv_CallCharMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jshort JNIEnv_CallShortMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jint JNIEnv_CallIntMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jlong JNIEnv_CallLongMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jfloat JNIEnv_CallFloatMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
jdouble JNIEnv_CallDoubleMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);
void JNIEnv_CallVoidMethod(JNIEnv* env, jobject obj, jmethodID id, jvalue* args);

jobject JNIEnv_CallNonvirtualObjectMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jboolean JNIEnv_CallNonvirtualBooleanMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jbyte JNIEnv_CallNonvirtualByteMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jchar JNIEnv_CallNonvirtualCharMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jshort JNIEnv_CallNonvirtualShortMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jint JNIEnv_CallNonvirtualIntMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jlong JNIEnv_CallNonvirtualLongMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jfloat JNIEnv_CallNonvirtualFloatMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
jdouble JNIEnv_CallNonvirtualDoubleMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);
void JNIEnv_CallNonvirtualVoidMethod(JNIEnv* env, jobject obj, jclass clazz, jmethodID id, jvalue* args);

jfieldID JNIEnv_GetFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

jobject JNIEnv_GetObjectField(JNIEnv* env, jobject obj, jfieldID id);
jboolean JNIEnv_GetBooleanField(JNIEnv* env, jobject obj, jfieldID id);
jbyte JNIEnv_GetByteField(JNIEnv* env, jobject obj, jfieldID id);
jchar JNIEnv_GetCharField(JNIEnv* env, jobject obj, jfieldID id);
jshort JNIEnv_GetShortField(JNIEnv* env, jobject obj, jfieldID id);
jint JNIEnv_GetIntField(JNIEnv* env, jobject obj, jfieldID id);
jlong JNIEnv_GetLongField(JNIEnv* env, jobject obj, jfieldID id);
jfloat JNIEnv_GetFloatField(JNIEnv* env, jobject obj, jfieldID id);
jdouble JNIEnv_GetDoubleField(JNIEnv* env, jobject obj, jfieldID id);

void JNIEnv_SetObjectField(JNIEnv* env, jobject obj, jfieldID id, jobject val);
void JNIEnv_SetBooleanField(JNIEnv* env, jobject obj, jfieldID id, jboolean val);
void JNIEnv_SetByteField(JNIEnv* env, jobject obj, jfieldID id, jbyte val);
void JNIEnv_SetCharField(JNIEnv* env, jobject obj, jfieldID id, jchar val);
void JNIEnv_SetShortField(JNIEnv* env, jobject obj, jfieldID id, jshort val);
void JNIEnv_SetIntField(JNIEnv* env, jobject obj, jfieldID id, jint val);
void JNIEnv_SetLongField(JNIEnv* env, jobject obj, jfieldID id, jlong val);
void JNIEnv_SetFloatField(JNIEnv* env, jobject obj, jfieldID id, jfloat val);
void JNIEnv_SetDoubleField(JNIEnv* env, jobject obj, jfieldID id, jdouble val);

jmethodID JNIEnv_GetStaticMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

jobject JNIEnv_CallStaticObjectMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jboolean JNIEnv_CallStaticBooleanMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jbyte JNIEnv_CallStaticByteMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jchar JNIEnv_CallStaticCharMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jshort JNIEnv_CallStaticShortMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jint JNIEnv_CallStaticIntMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jlong JNIEnv_CallStaticLongMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jfloat JNIEnv_CallStaticFloatMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
jdouble JNIEnv_CallStaticDoubleMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);
void JNIEnv_CallStaticVoidMethod(JNIEnv* env, jclass clazz, jmethodID id, jvalue* args);

jfieldID JNIEnv_GetStaticFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

jobject JNIEnv_GetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID id);
jboolean JNIEnv_GetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID id);
jbyte JNIEnv_GetStaticByteField(JNIEnv* env, jclass clazz, jfieldID id);
jchar JNIEnv_GetStaticCharField(JNIEnv* env, jclass clazz, jfieldID id);
jshort JNIEnv_GetStaticShortField(JNIEnv* env, jclass clazz, jfieldID id);
jint JNIEnv_GetStaticIntField(JNIEnv* env, jclass clazz, jfieldID id);
jlong JNIEnv_GetStaticLongField(JNIEnv* env, jclass clazz, jfieldID id);
jfloat JNIEnv_GetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID id);
jdouble JNIEnv_GetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID id);

void JNIEnv_SetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID id, jobject val);
void JNIEnv_SetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID id, jboolean val);
void JNIEnv_SetStaticByteField(JNIEnv* env, jclass clazz, jfieldID id, jbyte val);
void JNIEnv_SetStaticCharField(JNIEnv* env, jclass clazz, jfieldID id, jchar val);
void JNIEnv_SetStaticShortField(JNIEnv* env, jclass clazz, jfieldID id, jshort val);
void JNIEnv_SetStaticIntField(JNIEnv* env, jclass clazz, jfieldID id, jint val);
void JNIEnv_SetStaticLongField(JNIEnv* env, jclass clazz, jfieldID id, jlong val);
void JNIEnv_SetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID id, jfloat val);
void JNIEnv_SetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID id, jdouble val);

jstring JNIEnv_NewString(JNIEnv* env, jchar* buf, jsize bufLen);
jsize JNIEnv_GetStringLength(JNIEnv* env, jstring str);
const jchar* JNIEnv_GetStringChars(JNIEnv* env, jstring str, jboolean* isCopy);
void JNIEnv_ReleaseStringChars(JNIEnv* env, jstring str, jchar* chars);
jstring JNIEnv_NewStringUTF(JNIEnv* env, const char* str);
jsize JNIEnv_GetStringUTFLength(JNIEnv* env, jstring str);
const char* JNIEnv_GetStringUTFChars(JNIEnv* env, jstring str, jboolean* isCopy);
void JNIEnv_ReleaseStringUTFChars(JNIEnv* env, jstring str, const char* utf);
jsize JNIEnv_GetArrayLength(JNIEnv* env, jarray arr);
jobjectArray JNIEnv_NewObjectArray(JNIEnv* env, jsize length, jclass clazz, jobject obj);
jobject JNIEnv_GetObjectArrayElement(JNIEnv* env, jobjectArray arr, jsize index);
void JNIEnv_SetObjectArrayElement(JNIEnv* env, jobjectArray arr, jsize index, jobject obj);

jbooleanArray JNIEnv_NewBooleanArray(JNIEnv* env, jsize length);
jbyteArray JNIEnv_NewByteArray(JNIEnv* env, jsize length);
jcharArray JNIEnv_NewCharArray(JNIEnv* env, jsize length);
jshortArray JNIEnv_NewShortArray(JNIEnv* env, jsize length);
jintArray JNIEnv_NewIntArray(JNIEnv* env, jsize length);
jlongArray JNIEnv_NewLongArray(JNIEnv* env, jsize length);
jfloatArray JNIEnv_NewFloatArray(JNIEnv* env, jsize length);
jdoubleArray JNIEnv_NewDoubleArray(JNIEnv* env, jsize length);

jboolean* JNIEnv_GetBooleanArrayElements(JNIEnv* env, jbooleanArray arr, jboolean* isCopy);
jbyte* JNIEnv_GetByteArrayElements(JNIEnv* env, jbyteArray arr, jboolean* isCopy);
jchar* JNIEnv_GetCharArrayElements(JNIEnv* env, jcharArray arr, jboolean* isCopy);
jshort* JNIEnv_GetShortArrayElements(JNIEnv* env, jshortArray arr, jboolean* isCopy);
jint* JNIEnv_GetIntArrayElements(JNIEnv* env, jintArray arr, jboolean* isCopy);
jlong* JNIEnv_GetLongArrayElements(JNIEnv* env, jlongArray arr, jboolean* isCopy);
jfloat* JNIEnv_GetFloatArrayElements(JNIEnv* env, jfloatArray arr, jboolean* isCopy);
jdouble* JNIEnv_GetDoubleArrayElements(JNIEnv* env, jdoubleArray arr, jboolean* isCopy);

void JNIEnv_ReleaseBooleanArrayElements(JNIEnv* env, jbooleanArray arr, jboolean* elems, jint mode);
void JNIEnv_ReleaseByteArrayElements(JNIEnv* env, jbyteArray arr, jbyte* elems, jint mode);
void JNIEnv_ReleaseCharArrayElements(JNIEnv* env, jcharArray arr, jchar* elems, jint mode);
void JNIEnv_ReleaseShortArrayElements(JNIEnv* env, jshortArray arr, jshort* elems, jint mode);
void JNIEnv_ReleaseIntArrayElements(JNIEnv* env, jintArray arr, jint* elems, jint mode);
void JNIEnv_ReleaseLongArrayElements(JNIEnv* env, jlongArray arr, jlong* elems, jint mode);
void JNIEnv_ReleaseFloatArrayElements(JNIEnv* env, jfloatArray arr, jfloat* elems, jint mode);
void JNIEnv_ReleaseDoubleArrayElements(JNIEnv* env, jdoubleArray arr, jdouble* elems, jint mode);

void JNIEnv_GetBooleanArrayRegion(JNIEnv* env, jbooleanArray arr, jsize start, jsize length, jboolean* buf);
void JNIEnv_GetByteArrayRegion(JNIEnv* env, jbyteArray arr, jsize start, jsize length, jbyte* buf);
void JNIEnv_GetCharArrayRegion(JNIEnv* env, jcharArray arr, jsize start, jsize length, jchar* buf);
void JNIEnv_GetShortArrayRegion(JNIEnv* env, jshortArray arr, jsize start, jsize length, jshort* buf);
void JNIEnv_GetIntArrayRegion(JNIEnv* env, jintArray arr, jsize start, jsize length, jint* buf);
void JNIEnv_GetLongArrayRegion(JNIEnv* env, jlongArray arr, jsize start, jsize length, jlong* buf);
void JNIEnv_GetFloatArrayRegion(JNIEnv* env, jfloatArray arr, jsize start, jsize length, jfloat* buf);
void JNIEnv_GetDoubleArrayRegion(JNIEnv* env, jdoubleArray arr, jsize start, jsize length, jdouble* buf);

void JNIEnv_SetBooleanArrayRegion(JNIEnv* env, jbooleanArray arr, jsize start, jsize length, jboolean* buf);
void JNIEnv_SetByteArrayRegion(JNIEnv* env, jbyteArray arr, jsize start, jsize length, jbyte* buf);
void JNIEnv_SetCharArrayRegion(JNIEnv* env, jcharArray arr, jsize start, jsize length, jchar* buf);
void JNIEnv_SetShortArrayRegion(JNIEnv* env, jshortArray arr, jsize start, jsize length, jshort* buf);
void JNIEnv_SetIntArrayRegion(JNIEnv* env, jintArray arr, jsize start, jsize length, jint* buf);
void JNIEnv_SetLongArrayRegion(JNIEnv* env, jlongArray arr, jsize start, jsize length, jlong* buf);
void JNIEnv_SetFloatArrayRegion(JNIEnv* env, jfloatArray arr, jsize start, jsize length, jfloat* buf);
void JNIEnv_SetDoubleArrayRegion(JNIEnv* env, jdoubleArray arr, jsize start, jsize length, jdouble* buf);

jint JNIEnv_RegisterNatives(JNIEnv* env, jclass clazz, JNINativeMethod* methods, jint nMethods);
jint JNIEnv_UnregisterNatives(JNIEnv* env, jclass clazz);
jint JNIEnv_MonitorEnter(JNIEnv* env, jobject obj);
jint JNIEnv_MonitorExit(JNIEnv* env, jobject obj);
jint JNIEnv_GetJavaVM(JNIEnv* env, JavaVM** p_vm);

void JNIEnv_GetStringRegion(JNIEnv* env, jstring str, jsize start, jsize length, jchar* buf);
void JNIEnv_GetStringUTFRegion(JNIEnv* env, jstring str, jsize start, jsize length, char* buf);

void* JNIEnv_GetPrimitiveArrayCritical(JNIEnv* env, jarray arr, jboolean* isCopy);
void JNIEnv_ReleasePrimitiveArrayCritical(JNIEnv* env, jarray arr, void* carray, jint mode);

const jchar* JNIEnv_GetStringCritical(JNIEnv* env, jstring str, jboolean* isCopy);
void JNIEnv_ReleaseStringCritical(JNIEnv* env, jstring str, jchar* carray);

jweak JNIEnv_NewWeakGlobalRef(JNIEnv* env, jobject obj);
void JNIEnv_DeleteWeakGlobalRef(JNIEnv* env, jweak obj);

jboolean JNIEnv_ExceptionCheck(JNIEnv* env);

jobject JNIEnv_NewDirectByteBuffer(JNIEnv* env, void* buf, jlong capacity);
void* JNIEnv_GetDirectBufferAddress(JNIEnv* env, jobject buf);
jlong JNIEnv_GetDirectBufferCapacity(JNIEnv* env, jobject buf);

jobjectRefType JNIEnv_GetObjectRefType(JNIEnv* env, jobject obj);

#endif // JNI_CALL_H
