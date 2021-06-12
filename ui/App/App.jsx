import React, {useCallback, useState, useEffect} from 'react';

import user from "../api/resources/user";
import Login from "./views/Login";
import {Redirect, Route, Switch, useHistory} from "react-router";
import Controls from "./views/Controls";
import {BrowserRouter} from "react-router-dom";
import Logs from "./views/Logs";
import Saves from "./views/Saves/Saves";
import Layout from "./components/Layout";
import server from "../api/resources/server";
import Mods from "./views/Mods/Mods";
import UserManagement from "./views/UserManagement/UserManagment";
import ServerSettings from "./views/ServerSettings";
import GameSettings from "./views/GameSettings";
import Console from "./views/Console";
import Help from "./views/Help";
import socket from "../api/socket";
import {Flash} from "./components/Flash";
import MapGenerator from "./views/MapGenerator/MapGenerator";


const App = () => {

    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [serverStatus, setServerStatus] = useState(null);
    const history = useHistory();

    const updateServerStatus = async () => {
        const status = await server.status();
        if (status) {
            setServerStatus(status)
        }
    }

    const handleAuthenticationStatus = useCallback(async (status) => {
        if (status?.username) {
            setIsAuthenticated(true);
            await updateServerStatus()
            socket.emit('server status subscribe');
            socket.on('server_status', updateServerStatus);
        }
    },[]);

    const handleLogout = useCallback(async () => {
        const loggedOut = await user.logout();
        if (loggedOut) {
            setIsAuthenticated(false);
        }
    }, []);

    const ProtectedRoute = ({component: Component, ...rest}) => (
        <Route {...rest} render={(props) => (
            isAuthenticated && Component
                ? <Component serverStatus={serverStatus} updateServerStatus={updateServerStatus} {...props} />
                : <Redirect to={{
                    pathname: '/login',
                    state: {from: props.location}
                }}/>
        )}/>
    );

    useEffect(() => {
        (async () => {
            updateServerStatus()
        })();
    }, []);

    return (
        <BrowserRouter basename="/">
            <Switch>
                <Route path="/login" render={() => (<Login handleLogin={handleAuthenticationStatus}/>)}/>

                <Layout handleLogout={handleLogout} serverStatus={serverStatus}>
                    <ProtectedRoute exact path="/" component={Controls}/>
                    <ProtectedRoute path="/saves" component={Saves}/>
                    <ProtectedRoute path="/map-generator" component={MapGenerator}/>
                    <ProtectedRoute path="/mods" component={Mods}/>
                    <ProtectedRoute path="/server-settings" component={ServerSettings}/>
                    <ProtectedRoute path="/game-settings" component={GameSettings}/>
                    <ProtectedRoute path="/console" component={Console}/>
                    <ProtectedRoute path="/logs" component={Logs}/>
                    <ProtectedRoute path="/user-management" component={UserManagement}/>
                    <ProtectedRoute path="/help" component={Help}/>
                    <Flash/>
                </Layout>
            </Switch>
        </BrowserRouter>
    );
}

export default App;
