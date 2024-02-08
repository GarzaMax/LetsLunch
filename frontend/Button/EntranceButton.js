import React from "react";
import ReactDOM from "react-dom";
import '../Registration/Registration.css'
import {Link, Navigate} from "react-router-dom";

class EntranceButton extends React.Component {
    render() {
        return (<Link className='login-button' to={'/'}>Войти</Link>);
    }
}
export default EntranceButton;