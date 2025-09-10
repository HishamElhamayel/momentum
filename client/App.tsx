import { Container } from 'app/UI/Container';
import { AuthNavigator } from 'app/navigator/AuthNavigator';
import { StatusBar } from 'expo-status-bar';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import './global.css';

export default function App() {
  return (
    <>
      <StatusBar style="auto" />
      <SafeAreaProvider>
        <Container>
          <AuthNavigator />
        </Container>
      </SafeAreaProvider>
    </>
  );
}
