import React, {useState, useEffect, useContext} from "react";
import './Visit.css';
import axios from "axios";
import NavBar from "../NavBar/NavBar";
import CardVisit from "./CardVisit";
import tokenContext from "../tokenContext";

function Visit() {
    const {token, setToken} = useContext(tokenContext)
    const config = {headers: { Authorization: `Bearer ${token}` }}
    const [searchBarActive, setSearchBarActive] = useState(false);
    const [visits, setVisit] = useState([{
        url:"https://voyagist.ru/wp-content/uploads/2017/09/pekarni-sankt-peterburga-9.jpg",
        name:"Люди Любят",
        description:"Описания нет, придумайте сами"}]);

    useEffect(() => {
            fetchVisit()
        }, []
    )

    async function remove(id) {
        console.log('http://localhost:3333/api/meeting/update/' + id)
       await axios.delete('http://localhost:3333/api/meeting/update/' + id,
            config);
       fetchVisit()
    }

    async function fetchVisit() {
        const response = await axios.get('http://localhost:3333/api/meeting/find',
            config);
        console.log(response.data)
        setVisit(response.data)
    }

    return (
        <div className="account">
            <NavBar list={[
                {id: 0, link: "/mainpage/", description: "На главную"},
            ]}/>

            <h1 style={{marginTop: "80px"}}>Встречи</h1>
            <div className="list" style={{marginTop: "-5px", marginLeft: "20px"}}>
                {visits.map(card =>
                    <CardVisit url={"https://prorisuem.ru/foto/8812/narisovat_kafe_30.webp"} name={card.title} remove={remove} id={card.id}>{ }</CardVisit>)}
            </div>
        </div>
    );
}

export default Visit;





