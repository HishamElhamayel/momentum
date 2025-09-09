import { StatusBar } from 'expo-status-bar';

import { Container } from 'app/UI/Container';
import Register from 'app/screen/Register';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import './global.css';

export default function App() {
  return (
    <>
      <SafeAreaProvider>
        <Container>
          <StatusBar style="auto" />
          <Register />
        </Container>
      </SafeAreaProvider>
    </>
  );
}
