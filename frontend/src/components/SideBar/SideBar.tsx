import { Sidebar } from "flowbite-react";
import type { FC } from "react";
import Data from "./Data";
import Operations from "./Operations";

const SideBar: FC = function () {
    return (
        <Sidebar aria-label="Sidebar with multi-level dropdown example">
            <div>
                <Sidebar.ItemGroup>
                    <Data />
                    <Operations />
                </Sidebar.ItemGroup>
            </div>
        </Sidebar >
    );
};

export default SideBar;