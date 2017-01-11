package android

import "unsafe"

// typedef union jvalue {
//     jboolean    z;
//     jbyte       b;
//     jchar       c;
//     jshort      s;
//     jint        i;
//     jlong       j;
//     jfloat      f;
//     jdouble     d;
//     jobject     l;
// } jvalue;

func JbooleanV(z bool) Jvalue {
	if z {
		var b = JNITrue
		return *(*Jvalue)(unsafe.Pointer(&b))
	}
	var b = JNIFalse
	return *(*Jvalue)(unsafe.Pointer(&b))
}

func JbyteV(b byte) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&b))
}

func JcharV(c uint16) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&c))
}

func JshortV(s int16) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&s))
}

func JintV(i int32) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&i))
}

func JlongV(j int32) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&j))
}

func JfloatV(f float32) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&f))
}

func JdoubleV(d float64) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&d))
}

func JobjectV(l Jobject) Jvalue {
	return *(*Jvalue)(unsafe.Pointer(&l))
}
