package android

import "errors"

type SoftKeyboardState int32

const (
	SoftKeyboardHidden SoftKeyboardState = iota
	SoftKeyboardVisible
)

// SetSoftKeyboardState allows to toggle Android virtual keyboard using JNI calls into JavaVM.
func (a *NativeActivity) SetSoftKeyboardState(state SoftKeyboardState) error {
	return a.JNICall(func(env *JNIEnv, activity Jobject, activityClass, contextClass *Jclass) error {
		// Context.INPUT_METHOD_SERVICE
		inputMethodService := JNIEnvGetStaticObjectField(env, contextClass,
			JNIEnvGetStaticFieldID(env, contextClass,
				s("INPUT_METHOD_SERVICE"), JClassString.Spec().Sig()),
		)
		if inputMethodService == nil {
			return errors.New("failed to get INPUT_METHOD_SERVICE")
		}

		// getSystemService(Context.INPUT_METHOD_SERVICE)
		getSystemService := JNIEnvGetMethodID(env, activityClass,
			s("getSystemService"), JNIMethodSig(JClassObject.Spec(), JClassString.Spec()))
		inputMethodManager := JNIEnvCallObjectMethod(env, activity, getSystemService, []Jvalue{
			JobjectV(inputMethodService),
		})
		if inputMethodManager == nil {
			return errors.New("failed to run getSystemService()")
		}

		// getWindow().getDecorView()
		getWindowMethod := JNIEnvGetMethodID(env, activityClass,
			s("getWindow"), JNIMethodSig(JClassWindow.Spec()))
		getDecorViewMethod := JNIEnvGetMethodID(env, JNIEnvFindClass(env, JClassWindow.Name()),
			s("getDecorView"), JNIMethodSig(JClassView.Spec()))
		window := JNIEnvCallObjectMethod(env, activity, getWindowMethod, nil)
		if window == nil {
			return errors.New("failed to run getWindow()")
		}
		decorView := JNIEnvCallObjectMethod(env, window, getDecorViewMethod, nil)
		if decorView == nil {
			return errors.New("failed to run getDecorView()")
		}

		switch state {
		case SoftKeyboardHidden:
			const flags = 0

			// decorView.getWindowToken()
			getWindowTokenMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassView.Name()),
				s("getWindowToken"), JNIMethodSig(JClassIBinder.Spec()))
			binder := JNIEnvCallObjectMethod(env, decorView, getWindowTokenMethod, nil)
			if binder == nil {
				return errors.New("failed to run getWindowToken()")
			}

			// inputMethodManager.hideSoftInputFromWindow(...)
			hideSoftInputFromWindowMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassInputMethodManager.Name()),
				s("hideSoftInputFromWindow"), JNIMethodSig(JBoolean.Spec(), JClassIBinder.Spec(), JInt.Spec()))
			result := JNIEnvCallBooleanMethod(env, inputMethodManager,
				hideSoftInputFromWindowMethod, []Jvalue{
					JobjectV(binder), JintV(flags),
				})
			if result == JNIFalse {
				return errors.New("failed to run hideSoftInputFromWindow()")
			}

		case SoftKeyboardVisible:
			const flags = 0
			// inputMethodManager.showSoftInput(...)
			showSoftInputMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassInputMethodManager.Name()),
				s("showSoftInput"), JNIMethodSig(JBoolean.Spec(), JClassView.Spec(), JInt.Spec()))
			result := JNIEnvCallBooleanMethod(env, inputMethodManager,
				showSoftInputMethod, []Jvalue{
					JobjectV(decorView), JintV(flags),
				})
			if result == JNIFalse {
				return errors.New("failed to run showSoftInput()")
			}
		}
		return nil
	})
}
