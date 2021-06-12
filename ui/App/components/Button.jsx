import React from "react";
import {faSpinner} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";

const Button = ({ children, type, onClick, isSubmit, className, size, isLoading, isDisabled = false }) => {

    let color = '';
    let padding = '';

    switch (type) {
        case 'success':
            color = `bg-green ${isDisabled || isLoading ? "" : "hover:glow-green hover:bg-green-light" }`;
            break;
        case 'danger':
            color = `bg-red ${isDisabled || isLoading ? "" : "hover:glow-red hover:bg-red-light"}`;
            break;
        default:
            color = `bg-gray-light ${isDisabled || isLoading ? "" : "hover:glow-orange hover:bg-orange"}`
    }

    switch (size) {
        case 'sm':
            padding = 'py-1 px-2';
            break;
        case 'none':
            padding = "";
            break;
        default:
            padding = 'py-2 px-4'
    }

    return (
        <button onClick={onClick} disabled={isDisabled || isLoading} className={`${className ? className: ""} ${isDisabled || isLoading ? "bg-opacity-50 cursor-not-allowed" : ""} ${padding} ${color} inline-block accentuated text-black font-bold`}
        type={isSubmit ? 'submit' : 'button'}>
            {children} { isLoading && <FontAwesomeIcon icon={faSpinner} spin={true}/>}
        </button>
    );
}

export default Button;