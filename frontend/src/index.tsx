import { createContext } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import UserStore from './store/UserStore';
import { Flowbite } from "flowbite-react";
import theme from "./flowbite-theme";
import { ToastContainer } from 'react-toastify';
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";


const queryClient = new QueryClient()

interface State {
  userStore: UserStore,
}

const userStore = new UserStore();

export const Context = createContext<State>({
  userStore,
})

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <div>
    <QueryClientProvider client={queryClient}>
      <ToastContainer />
      <Flowbite theme={{ theme }}>
        <Context.Provider value={{
          userStore
        }}>
          <App />
        </Context.Provider>
      </Flowbite>
    </QueryClientProvider>
  </div>
);
