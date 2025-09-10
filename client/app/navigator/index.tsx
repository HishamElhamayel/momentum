import Home from 'app/screen/user/Home';
import { useAuthStore } from 'app/store/AuthStore';
import { StyleSheet } from 'react-native';

const Navigator = () => {
  const { user } = useAuthStore();

  return <Home />;
};

export default Navigator;

const styles = StyleSheet.create({});
