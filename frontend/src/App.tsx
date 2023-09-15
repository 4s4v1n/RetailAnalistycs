import { FC, useContext, useEffect } from 'react';
import { BrowserRouter } from 'react-router-dom';
import AppRouter from './components/AppRouter/AppRouter';
import NavBar from './components/NavBar/NavBar';
import SideBar from './components/SideBar/SideBar';
import { Context } from '.';
import { observer } from "mobx-react-lite";

const App: FC = () => {
  const { userStore } = useContext(Context)

  useEffect(() => {
    if (localStorage.getItem("token")) {
      userStore.checkAuth();
    }
  }, [])

  if (!userStore.isAuth) {
    return (
      <BrowserRouter>
        <NavBar />
        <AppRouter />
      </BrowserRouter>
    );
  }

  return (
    <BrowserRouter>
      <NavBar />
      <SideBar />
      <AppRouter />
    </BrowserRouter>
  );
}

export default observer(App);
