import React, {useContext} from "react";
import './account.css';
import NavBar from "../NavBar/NavBar";
import personContext from "../personContext";
function Account() {
    const {person, setPerson} = useContext(personContext)


    return(
        <div className="account">

            <NavBar list={[
                {id: 1, link: "/mainpage/", description: "На главную"},
            ]}/>

            <div className="book" id="book">
                <h1 className="heading">Профиль</h1>
                <div className="row">
                    <form className="form" action="">
                        <div className="inputBox">
                            <h3>Имя</h3>
                            <input type="text" placeholder="place name"/>
                            <h3>Дата рождения</h3>
                            <input  className="date" type="date" placeholder="birthday"/>
                        </div>
                        <div className="inputBox">
                            <h3>О себе</h3>
                            <input type="text" className="about" placeholder="place something about you"/>
                        </div>
                        <input type="submit" className="btn" value="Сохранить"/>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Account;