import { FC, useContext } from "react";
import { Route, Routes } from "react-router-dom";
import { authRoutes, publicRoutes } from "../../route";
import { Context } from "../..";

const AppRouter: FC = () => {
    const { userStore } = useContext(Context)
    return (
        <Routes>
            {userStore.isAuth && authRoutes.map(({ path, Component }) =>
                <Route key={path} path={path} Component={Component} />
            )}

            {publicRoutes.map(({ path, Component }) =>
                <Route key={path} path={path} Component={Component} />
            )}
        </Routes>
    )
}

export default AppRouter;