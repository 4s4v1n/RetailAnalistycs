import { FC, useContext } from "react"

import { Button, Navbar } from "flowbite-react";

import {
    HiOutlineArrowRight,
    HiLogin,
} from "react-icons/hi";
import { HOME_ROUTE, SIGIN_ROUTE } from "../../utils/const";
import { Context } from "../..";
import { notifySucces } from "../Notify/Notify";
import { useNavigate } from "react-router-dom";

const NavBar: FC = () => {
    const { userStore } = useContext(Context);
    const navigate = useNavigate();

    const handleClickExit = async () => {
        userStore.logout();
        notifySucces("success log out");
        navigate(HOME_ROUTE);
      };

    return (
        <Navbar fluid>
            <div className="w-full p-3 lg:px-5 lg:pl-3">
                <div className="flex items-center justify-between">
                    <div className="flex items-center">
                        <Navbar.Brand href="/">
                            <span className="self-center whitespace-nowrap text-2xl font-semibold dark:text-white">
                                Retail Analitics
                            </span>
                        </Navbar.Brand>
                    </div>
                    <div className="flex items-center gap-2">
                        {userStore.isAuth
                            ?
                            <Button
                                color="primary"
                                onClick={handleClickExit}
                            >
                                <HiLogin className="ml-2 h-5 w-5" />
                            </Button> :
                            <Button
                                color="primary"
                                href={SIGIN_ROUTE}
                            >
                                Sign in
                                <HiOutlineArrowRight className="ml-2 h-5 w-5" />
                            </Button>
                        }
                    </div>
                </div>

            </div>
        </Navbar >
    );
};

export default NavBar;