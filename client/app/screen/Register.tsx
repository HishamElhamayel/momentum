import FormInput from 'app/components/form/FormInput';
import AuthButton from 'app/UI/buttons/AuthButton';
import { KeyboardAvoidingView, Platform, ScrollView, Text, View } from 'react-native';

const Register = () => {
  return (
    <KeyboardAvoidingView
      className="flex-1"
      behavior="padding"
      keyboardVerticalOffset={Platform.OS === 'ios' ? 30 : 30}>
      <ScrollView
        keyboardShouldPersistTaps="handled"
        contentContainerStyle={{ flexGrow: 1 }}
        className="flex-1 px-4">
        <Text className="mt-3 text-center text-3xl font-bold text-gray-900">Register</Text>
        <Text className="mt-2 text-center text-gray-500">Create an account to get started</Text>
        <View className="my-5">
          <FormInput placeholder={`Email or Username`} />
          <FormInput placeholder="Password" />
        </View>
        <View className="mt-auto pb-6">
          <AuthButton className="">Register</AuthButton>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
};

export default Register;
