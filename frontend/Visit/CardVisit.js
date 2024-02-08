import Button from "../Button/Button";
import Modal from "../Modal/Modal";
import React, {useContext, useState} from "react";
import visitContext from "../visitContext";

const CardVisit = (props) => {
    const {meeting, setMeeting} = useContext(visitContext)
    const [modal, setModal] = useState(false);

    const fun = () => {
        meeting ?
            <Modal visible={modal} setVisible={setModal(true)}>
                <h1>Вы не можете присоединиться, у вас своя встреча!</h1>
            </Modal>
            :
            setMeeting({
                    url: "https://prorisuem.ru/foto/8812/narisovat_kafe_30.webp",
                    title: props.name
                })

        meeting ?
            <></>
            :  props.remove(props.id)
    }

    return (
        <div className="card">
            <img src={props.url} alt={props.name}
                 className="card_img"/>
            <div className="card_body">
                <h3 className="card_title">{props.name}</h3>
                <p className="card_desc">{props.children}</p>
                <Button onClick={() => fun()} className="event">Присоединиться</Button>
                <Modal visible={modal} setVisible={setModal}>
                    <h1>Вы не можете присоединиться у вас уже есть встреча!</h1>
                </Modal>

            </div>
        </div>
    );
};

export default CardVisit;