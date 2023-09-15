import { FC } from "react";
import { observer } from "mobx-react-lite";

const Home: FC = () => {
    return (
        <div>
            <div className="sm:ml-64 p-4 pt-20">
                <p className="text-7xl text-gray-900 dark:text-white">Home Page</p>
                <p className="text-6xl text-gray-900 dark:text-white">Telvina - backend | Adough - frontend</p>
                <img className="h-auto max-w-full" src="/img/main.png" alt="main view" />
            </div>
        </div>
    )
}

export default observer(Home);