import React from "react";

const Button = ({ children, type, onClick, isSubmit, className, size }) => {

    let color = '';
    let padding = '';

    switch (type) {
        case 'success':
            color = 'bg-green hover:glow-green hover:bg-green-light';
            break;
        case 'danger':
            color = 'bg-red hover:glow-red hover:bg-red-light';
            break;
        default:
            color = 'bg-gray-light hover:glow-orange hover:bg-orange'
    }

    switch (size) {
        case 'sm':
            padding = 'py-1 px-2';
            break;
        default:
            padding = 'py-2 px-4'
    }

    return (
        <button onClick={onClick} className={`${className ? className: null} ${padding} ${color} inline-block accentuated text-black font-bold`}
        type={isSubmit ? 'submit' : 'button'}>
            {children}
        </button>
    );
}

export default Button;