import { NavigationProp, useNavigation } from '@react-navigation/native';
import { AuthStackParamList } from 'app/navigator/AuthNavigator';
import AuthButton from 'app/UI/buttons/AuthButton';
import FormInput from 'app/UI/form/FormInput';
import { Text, View } from 'react-native';

const LoginForm = () => {
  const navigation = useNavigation<NavigationProp<AuthStackParamList>>();
  return (
    <>
      <View className="my-5">
        <FormInput placeholder={`Email or Username`} />
        <FormInput placeholder="Password" />
      </View>

      <View className="mb-2 mt-auto pb-6">
        <AuthButton className="bg-surface">Login</AuthButton>
        <Text className="my-4 text-center text-gray-500">
          Don&apos;t have an account?{' '}
          <Text className="font-bold text-accent" onPress={() => navigation.navigate('Register')}>
            Register
          </Text>
        </Text>
      </View>
    </>
  );
};

export default LoginForm;
