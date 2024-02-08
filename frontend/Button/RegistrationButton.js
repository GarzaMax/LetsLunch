import React from "react";
import ReactDOM from "react-dom";
import  "../Entrance/entranceStyle.css";

class RegistrationButton extends React.Component {
    onclick () {
        window.location.assign('http://localhost:3000/registration');
    }
    render() {
        return (<button className="register-button" onClick={(e) => this.onclick(e.preventDefault())}>Зарегистрироваться</button>);
    }
}
export default RegistrationButton;