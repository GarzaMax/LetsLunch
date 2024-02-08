import React from "react";
import ReactDOM from "react-dom";
import  "../components/formInput.css";

class ExitButton extends React.Component {
    onclick () {
        window.location.assign('http://localhost:3000/');
    }
    render() {
        return (<button  className="exit-button" onClick={(e) => this.onclick(e.preventDefault())}>Выйти</button>);
    }
}
export default ExitButton;