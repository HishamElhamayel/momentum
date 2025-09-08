import { TextInput } from 'react-native';

const FormInput = ({ placeholder }: { placeholder: string }) => {
  return <TextInput placeholder={placeholder} className="rounded-md border border-gray-300 p-2" />;
};

export default FormInput;
