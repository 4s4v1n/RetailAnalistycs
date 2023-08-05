import { toast } from 'react-toastify';
import "react-toastify/dist/ReactToastify.css";

export const notifySucces = (message: string) => {
    toast.success(message, {
      position: toast.POSITION.BOTTOM_RIGHT
    });
  };

export const notifyError = (message: string) => {
    toast.error(message, {
      position: toast.POSITION.BOTTOM_RIGHT
    });
  };