import Button from "../Button/Button";
import Modal from "../Modal/Modal";
import Form from "../Modal/Form";
import './Card.css'
import React, {useState} from "react";

const Card = ({url, name, children, id}) => {
    const [modal, setModal] = useState(false);


    return (
        <div className="card">
        <img src={url} alt={name}
             className="card_img"/>
        <div className="card_body">
            <h3 className="card_title">{name}</h3>
            <p className="card_desc">{children}</p>
            <Button onClick={() => setModal(true)} className="event">Назначить встречу</Button>
            <Modal visible={modal} setVisible={setModal}>
                <Form setVisible={setModal} url={url} name={name} id={id}/>
            </Modal>

        </div>
    </div>
    );
};

export default Card;