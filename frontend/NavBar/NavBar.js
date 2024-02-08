import React from "react";
import './NavBar.css'
import {Link} from "react-router-dom";

const NavBar = ({list}) => {
    return (
            <header>
                <Link to="#" className="logo"><h1>L</h1>et's Lunch</Link>
                <div className="navbar">
                    {list.map(element => <Link  key={element.id} to={element.link}>{element.description}</Link>)}
                </div>
            </header>
        )
}

export default NavBar;