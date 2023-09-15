import { Sidebar } from "flowbite-react";
import type { FC } from "react";
import {
    HiChartPie,
} from "react-icons/hi";
import { DataLabel } from "../../store/Data";


const Data: FC = function () {

    return (
        <Sidebar.ItemGroup>
            <Sidebar.Collapse
                icon={HiChartPie}
                label="Data"
            >
                {DataLabel.map((item, i) => (
                    <Sidebar.Item key={i} href={item.link}>
                        {item.name}
                    </Sidebar.Item>
                ))}
            </Sidebar.Collapse>
        </Sidebar.ItemGroup>
    );
};

export default Data;