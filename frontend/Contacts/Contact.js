import React from "react";


import './Contacts.css';
import NavBar from "../NavBar/NavBar";
function Contact() {
    return (
        <div className="contact">

            <NavBar list={[
                {id: 0, link: "/mainpage/", description: "На главную"},
            ]}/>

            <div className="info" style={{marginTop: '80px', fontSize: "2em", marginLeft: "10px", color: "#542200"}}>
            <h2>Над приложением работали:</h2>
                <h3>Акимцов Максим</h3>
                <h3>Бородина Ирина</h3>
                <h3>Бровкина София</h3>
                <h3>Котельникова Елизавета</h3>
                <h3>Хейфец Михаил</h3>
            </div>
        </div>
    )
}

export default Contact;