import React from "react";

const Input = ({
                   name,
                   inputRef,
                   placeholder = null,
                   isInline = false,
                   type = "text",
                   defaultValue = undefined,
                   hasAutoComplete = true,
                   onKeyDown = () => null,
                   onChange = () => null,
                   min = null,
                   value = undefined,
                   disabled = false,
                   size = 'default'
               }) => {

    let padding;

    switch (size) {
        case 'sm':
            padding = 'py-1 px-2';
            break;
        case 'none':
            padding = "";
            break;
        default:
            padding = 'py-2 px-3'
    }

    return (
        <input
            className={
                `shadow rounded-sm appearance-none border border-gray-light placeholder-white ${padding} `
                + (isInline ? 'w-48 inline-block ' : 'w-full block ')
                + (disabled ? 'bg-gray-dark text-gray-light cursor-not-allowed' : 'bg-gray-light text-black')
            }
            placeholder={placeholder}
            ref={inputRef}
            name={name}
            id={name}
            type={type}
            onKeyDown={onKeyDown}
            onChange={onChange}
            autoComplete={hasAutoComplete ? "on" : "off"}
            defaultValue={defaultValue}
            min={min}
            value={value}
            disabled={disabled}
        />
    )
}

export default Input;