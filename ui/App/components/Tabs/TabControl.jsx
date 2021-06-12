import React, {useState} from "react";
import TabTitle from "./TabTitle";
import Panel from "../Panel";

const TabControl = ({children, actions = null, title= null}) => {
    const [selectedTab, setSelectedTab] = useState(0)

    return (
        <div className="mb-6">
            <div className="px-4 pt-3">
                {children.map((item, index) => (
                    <TabTitle
                        key={index}
                        title={item.props.title}
                        index={index}
                        isActive={index === selectedTab}
                        setSelectedTab={setSelectedTab}
                    />
                ))}
            </div>
            <Panel
                className="z-10 relative"
                content={children[selectedTab]}
                title={title}
                actions={actions}
            />
        </div>
    )
}

export default TabControl