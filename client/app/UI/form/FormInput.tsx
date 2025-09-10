import { TextInput, TextInputProps } from 'react-native';

const FormInput = (props: TextInputProps) => {
  return (
    <TextInput
      className="my-3 rounded-lg border border-gray-300 bg-gray-100 px-3 py-5 focus:border-accent focus:ring-accent"
      {...props}
    />
  );
};

export default FormInput;
