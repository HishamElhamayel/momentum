import { ReactNode } from 'react';
import { Pressable, PressableProps, Text } from 'react-native';

interface Props extends PressableProps {
  children: ReactNode;
  className?: string;
}

const WhiteButton = ({ children, className, ...props }: Props) => {
  return (
    <Pressable
      className={`w-full items-center rounded-lg bg-gray-300 py-4 active:bg-gray-400 ${className}`}
      {...props}>
      <Text className="">{children}</Text>
    </Pressable>
  );
};

export default WhiteButton;
