import RegisterForm from 'app/components/form/RegisterForm';
import { KeyboardAvoidingView, Platform, Text, View } from 'react-native';

const Register = () => {
  return (
    <KeyboardAvoidingView
      className="flex-1"
      behavior={Platform.OS === 'ios' ? 'padding' : 'padding'}
      keyboardVerticalOffset={Platform.OS === 'ios' ? 10 : 0}>
      <View className="flex-1">
        <Text className="mt-3 text-center text-3xl font-bold text-gray-900">Register</Text>
        <Text className="mt-2 text-center text-gray-500">Create an account to get started</Text>
        <RegisterForm />
      </View>
    </KeyboardAvoidingView>
  );
};

export default Register;
