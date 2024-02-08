import React, {useState, useEffect, useContext} from "react";
import './styleMain.css';
import Card from "./Card";
import MyCard from "./MyCard";
import axios from "axios";
import NavBar from "../NavBar/NavBar";
import tokenContext from "../tokenContext";
import visitContext from "../visitContext";
import personContext from "../personContext";

function MainPage() {
    const {token, setToken} = useContext(tokenContext)
    const [searchBarActive, setSearchBarActive] = useState(false);
    const {meeting, setMeeting} = useContext(visitContext)
    const [cards, setCards] = useState([]);
    const {person, setPerson} = useContext(personContext)
    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };

    useEffect(() => {
        fetchCard()
    }, []);


    async function fetchCard() {
        try {
            console.log(person)
            const response = await axios.get('http://localhost:3333/api/gatheringPlace/find', config);
            setCards(response.data)
            console.log(response.data)
        } catch (error) {
        }
    }

    return (
        <div className="Main">

            <NavBar list={[
                {id: 0, link: "/", description: "Выход"},
                {id: 1, link: "/mainpage/account/", description: "Аккаунт"},
                {id: 2, link: "/mainpage/visit/", description: "Вcтречи"},
                {id: 3, link: "/mainpage/contact/", description: "Контакты"}
            ]}/>
            {meeting ?
                <div>
                    <h2>Моя встреча</h2>
                    <div  style={{display: 'flex', justifyContent: 'center'}}>
                    <MyCard url={meeting.url} name={meeting.title}>{meeting.description}</MyCard>
                    </div>
                    <h2>Куда отправляемся?</h2>
                </div>
                : <h2>Куда отправляемся?</h2>
            }

            <div className="list">
                {cards.map(card =>
                <Card url={card.description} name={card.title} id={card.id}>{"Уютное место для встреч"}</Card>)}
            </div>

        </div>
    );
}

export default MainPage;


