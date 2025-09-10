import { NavigationContainer } from '@react-navigation/native';
import { SafeAreaView } from 'react-native-safe-area-context';

export const Container = ({ children }: { children: React.ReactNode }) => {
  return (
    <NavigationContainer>
      <SafeAreaView className={styles.container}>{children}</SafeAreaView>
    </NavigationContainer>
  );
};

const styles = {
  container: 'flex flex-1 m-4',
};
