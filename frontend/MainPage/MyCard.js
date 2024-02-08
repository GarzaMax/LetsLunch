import Button from "../Button/Button";
import Modal from "../Modal/Modal";
import Form from "../Modal/Form";
import React, {useContext, useState} from "react";
import visitContext from "../visitContext";

const MyCard = ({url, name, children}) => {
    const [modal, setModal] = useState(false);
    const {meeting, setMeeting} = useContext(visitContext)


    return (
        <div className="card">
            <img src={url} alt={name}
                 className="card_img"/>
            <div className="card_body">
                <h3 className="card_title">{name}</h3>
                <p className="card_desc">{children}</p>
                <h1>{name}</h1>
                <Button onClick={() => setMeeting('')} className="event">Удалить</Button>
                <Modal visible={modal} setVisible={setModal}>
                    <Form setVisible={setModal}/>
                </Modal>

            </div>
        </div>
    );
};

export default MyCard;