import { Sidebar } from "flowbite-react";
import type { FC } from "react";
import {
    HiNewspaper,
} from "react-icons/hi";
import { FunctionLabel, ViewLabel } from "../../store/Data";


const Operations: FC = function () {
    return (
        <Sidebar.ItemGroup>
            <Sidebar.Collapse
                icon={HiNewspaper}
                label="Operations"
            >
                <Sidebar.Collapse
                    label="Views"
                >
                    {ViewLabel.map((item, i) => (
                        <Sidebar.Item key={i} href={item.link}>
                            {item.name}
                        </Sidebar.Item>
                    ))}
                </Sidebar.Collapse>
                <Sidebar.Collapse
                    label="Functions"
                >
                    {FunctionLabel.map((item, i) => (
                        <Sidebar.Item key={i} href={item.link}>
                            {item.name}
                        </Sidebar.Item>
                    ))}
                </Sidebar.Collapse>
            </Sidebar.Collapse>
        </Sidebar.ItemGroup>
    );
};

export default Operations;