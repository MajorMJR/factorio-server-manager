import React, {useEffect, useState} from "react";
import {NavLink} from "react-router-dom";
import Button from "./Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBars} from "@fortawesome/free-solid-svg-icons";

const Layout = ({children, handleLogout, serverStatus}) => {

    const [isNavCollapsed, setIsNavCollapsed] = useState(true);



    const Status = ({info}) => {

        let text = 'Unknown';
        let color = 'gray-light';

        if (info) {
            if (info.status === 'running') {
                text = 'Running';
                color = 'green';
            } else if (info.status === 'stopped') {
                text = 'Stopped';
                color = 'red';
            }
        }

        return (
            <div className={`bg-${color} accentuated rounded px-2 py-1 text-black`}>{text}</div>
        )
    }

    const Link = ({children, to, last}) => {
        return (
            <NavLink
                onClick={() => setIsNavCollapsed(true)}
                exact={true}
                to={to}
                activeClassName="bg-orange"
                className={`hover:glow-orange accentuated bg-gray-light hover:bg-orange text-black font-bold py-2 px-4 w-full block${last ? '' : ' mb-1'}`}
            >{children}</NavLink>)
    }

    return (
        <>
            {/*Sidebar*/}
            <div className="w-full md:w-88 md:fixed md:top-0 md:left-0 bg-gray-dark md:h-screen overflow-y-auto">
                <div className="py-4 px-2 accentuated">
                    <div className="mx-4 justify-between flex text-center">
                        <span className="text-dirty-white text-xl">Factorio Server Manager</span>
                        <button
                            className="md:hidden cursor-pointer text-white hover:text-dirty-white"
                            onClick={() => setIsNavCollapsed(!isNavCollapsed)}
                        >
                            <FontAwesomeIcon icon={faBars}/>
                        </button>
                    </div>
                </div>
                <div className={isNavCollapsed ? "hidden md:block" : "block"}>
                    <div className="py-4 px-2 accentuated">
                        <h1 className="text-dirty-white text-lg mb-2 mx-4">Server Status</h1>
                        <div className="mx-4 mb-4 text-center">
                            <Status info={serverStatus}/>
                        </div>
                    </div>
                    <div className="py-4 px-2 accentuated">
                        <h1 className="text-dirty-white text-lg mb-2 mx-4">Server Management</h1>
                        <div className="text-white text-center rounded-sm bg-black shadow-inner mx-4 p-1">
                            <Link to="/">Controls</Link>
                            <Link to="/saves">Saves</Link>
                            <Link to="/map-generator">Map Generator</Link>
                            <Link to="/mods">Mods</Link>
                            <Link to="/server-settings">Server Settings</Link>
                            <Link to="/game-settings">Game Settings</Link>
                            <Link to="/console">Console</Link>
                            <Link to="/logs" last={true}>Logs</Link>
                        </div>
                    </div>
                    <div className="py-4 px-2 accentuated">
                        <h1 className="text-dirty-white text-lg mb-2 mx-4">FSM Administration</h1>
                        <div className="text-white text-center rounded-sm bg-black shadow-inner mx-4 p-1">
                            <Link to="/user-management">Users</Link>
                            <Link to="/help" last={true}>Help</Link>
                        </div>
                    </div>
                    <div className="py-4 px-2 accentuated">
                        <div className="text-white text-center rounded-sm bg-black shadow-inner mx-4 p-1">
                            <Button type="danger" className="w-full" onClick={handleLogout}>Logout</Button>
                        </div>
                    </div>
                    <div className="accentuated-t accentuated-x md:block hidden"/>
                </div>
            </div>

            {/*Main*/}
            <div className="md:ml-88 min-h-screen">
                <div className="container mx-auto pt-16 px-6">
                    {children}
                </div>
            </div>
        </>
    );
}

export default Layout;