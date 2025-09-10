import { NavigationProp, useNavigation } from '@react-navigation/native';
import { AuthStackParamList } from 'app/navigator/AuthNavigator';
import AuthButton from 'app/UI/buttons/AuthButton';
import WhiteButton from 'app/UI/buttons/WhiteButton';
import FormInput from 'app/UI/form/FormInput';

import { View } from 'react-native';

const RegisterForm = () => {
  const navigation = useNavigation<NavigationProp<AuthStackParamList>>();
  return (
    <>
      <View className="my-5">
        <FormInput placeholder={`Name`} />
        <FormInput placeholder={`Email or Username`} />
        <FormInput placeholder="Password" />
      </View>
      <View className=" mt-auto pb-6 ">
        <AuthButton className="my-1 bg-surface">Register</AuthButton>
        <WhiteButton className="my-1" onPress={() => navigation.goBack()}>
          Back
        </WhiteButton>
      </View>
    </>
  );
};

export default RegisterForm;
