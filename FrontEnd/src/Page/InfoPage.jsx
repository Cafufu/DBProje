import React, { useEffect, useState } from 'react'
import { useCookies } from 'react-cookie';
import '../css/info.css'
import '../css/bill.css'
function InfoPage() {
    const [responseMessage, setResponseMessage] = useState('');
    const [allStorageData, setAllStorageData] = useState({});
    const [cookies, setCookie, getCookies] = useCookies(['customerData']);
    const [billInfo, setbillInfo] = useState({
        userId: cookies.customerData,
        billType: '1',
        billname: '',
        month: '',
        year: '',
        amount: '',
    });
    const [show, setShow] = useState({
        userId: cookies.customerData,
        billType: 0,
    });
    const handleInputLogin = (e) => {
        const { name, value } = e.target;
        setbillInfo(prevState => ({
            ...prevState,
            [name]: value,
        }));
    }

    const allFieldsFilled = () => {
        return (
            billInfo.billType &&
            billInfo.billname &&
            billInfo.month &&
            billInfo.year &&
            billInfo.amount
        );
    }

    const insertDB = () => {
        fetch('http://localhost:3000/insert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(billInfo),
        })
            .then(response => response.json())
            .then(data => {
                if (data === 1) {
                    setResponseMessage('Kayıt başarılı');
                    console.log("başarılı");
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    const showDB = (typeId) => {
        show.billType = typeId
        fetch('http://localhost:3000/show', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(show),
        })
            .then(response => response.json())
            .then(data => {
                if (data === 1) {
                    console.log(data);
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }


    return (
        <div>
            <div className="info-page">
                <div className="left-panel">
                    <h2>Yeni Fatura Ekle</h2>
                    <div style={{ textAlign: "center" }}>
                        {responseMessage}
                    </div>
                    <form className="bill-form">

                        <div className="form-group">
                            <label htmlFor="bill-type">Fatura Türü</label>
                            <select
                                id="bill-type"
                                name="billType"
                                value={billInfo.billType}
                                onChange={handleInputLogin}
                            >
                                <option value="1">Elektrik</option>
                                <option value="2">Doğalgaz</option>
                                <option value="3">Su</option>
                            </select>
                        </div>


                        <div className="form-group">
                            <label htmlFor="bill-name">Fatura Adı</label>
                            <input
                                type="text"
                                id="billname"
                                name="billname"
                                placeholder="Faturanın adı"
                                value={billInfo.billname}
                                onChange={handleInputLogin}
                            />
                        </div>


                        <div className="form-group row">
                            <div className="column">
                                <label htmlFor="month">Ay</label>
                                <input
                                    type="number"
                                    id="month"
                                    name="month"
                                    placeholder="1"
                                    min="1"
                                    max="12"
                                    value={billInfo.month}
                                    onChange={handleInputLogin}
                                />
                            </div>

                            <div className="column">
                                <label htmlFor="year">Yıl</label>
                                <input
                                    type="number"
                                    id="year"
                                    name="year"
                                    placeholder="2024"
                                    step="1"
                                    min="2000"
                                    max="2100"
                                    value={billInfo.year}
                                    onChange={handleInputLogin}
                                />
                            </div>
                        </div>

                        <div className="form-group">
                            <label htmlFor="amount">Fatura Tutarı</label>
                            <input
                                type="number"
                                id="amount"
                                name="amount"
                                placeholder="0.00"
                                step="0.01"
                                value={billInfo.amount}
                                onChange={handleInputLogin}
                            />
                        </div>

                        <button
                            className="btn-save"
                            type="button"
                            disabled={!allFieldsFilled()}
                            onClick={insertDB}
                        >Kaydet
                        </button>

                    </form>
                </div>

                <div className="right-panel">
                    <h2>Fatura Görüntüleme ve Analiz</h2>
                    <div className="button-group">
                        <button className="view-button" onClick={() => showDB(1)}>
                            Elektrik Faturasını Görüntüle
                        </button>
                        <button className="view-button" onClick={() => showDB(2)}>
                            Su Faturasını Görüntüle
                        </button>
                        <button className="view-button" onClick={() => showDB(3)}>
                            Doğalgaz Faturasını Görüntüle
                        </button>
                        <button className="analyze-button">Karbon Ayak İzini Hesapla</button>
                        <button className="analyze-button">Analiz Yap</button>
                    </div>
                </div>
            </div>
            <div className="output-panel">
                <h3>Çıktı Ekranı</h3>
                <div className="output-content">
                    <div className="cart-header">
                        <h1>Sepetim</h1>
                    </div>
                    {Object.keys(allStorageData).map((key) => (
                        allStorageData[key].map(item => (
                            <div className="cart-item">
                                <div className="cart-item-info">
                                    <div className="basket-info">
                                        <img src={item.object.ImageUrl}></img>
                                        <div className="basket-details">
                                            <div className="basket-name">{item.object.Brand}{item.object.Model}</div>
                                            <div className="basket-size">Beden: {item.size}</div>
                                        </div>
                                    </div>
                                </div>
                                <div className="cart-item-actions">
                                    <div className="price">{item.amount}₺</div>
                                    <button className="delete-btn" onClick={() => removeItemFromLocalStorage(item.Id)}>Sil</button>
                                </div>
                            </div>

                        ))
                    ))}
                    Henüz bir veri yok. Lütfen bir butona tıklayın.<br />
                </div>
            </div>
        </div>
    )
}

export default InfoPage