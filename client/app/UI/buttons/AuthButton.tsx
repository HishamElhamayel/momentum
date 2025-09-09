import { ReactNode } from 'react';
import { Pressable, PressableProps, Text } from 'react-native';

interface Props extends PressableProps {
  children: ReactNode;
  className?: string;
}

const AuthButton = ({ children, className }: Props) => {
  return (
    <Pressable
      className={`bg-surface active:bg-accent w-full items-center rounded-lg py-4 ${className}`}>
      <Text className="text-white">{children}</Text>
    </Pressable>
  );
};

export default AuthButton;
