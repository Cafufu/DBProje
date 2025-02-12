import React, { useEffect, useState } from 'react'
import { useCookies } from 'react-cookie';
import Header from '../components/header.jsx'
import '../css/info.css'
import '../css/bill.css'
function InfoPage() {
    const [responseMessage, setResponseMessage] = useState('');
    const [allStorageData, setAllStorageData] = useState([]);
    const [typeName, setTypeName] = useState("");
    const [carbonFootprint, setCarbonFootprint] = useState("");
    const [Analyze, setAnalyze] = useState("");
    const [cookies] = useCookies(['customerData']);
    const [billInfo, setbillInfo] = useState({
        userId: cookies.customerData,
        billType: '1',
        billname: '',
        month: '',
        year: '',
        amount: 0,
    });
    const [show, setShow] = useState({
        userId: cookies.customerData,
    });

    const handleInputLogin = (e) => {
        const { name, value } = e.target;

        if (name === 'amount') {
            const intValue = value === '' ? '' : parseInt(value, 10);
            setbillInfo(prevState => ({
                ...prevState,
                [name]: intValue,
            }));
        } else {
            setbillInfo(prevState => ({
                ...prevState,
                [name]: value,
            }));
        }
    };

    const deleteItem = (indexToDelete) => {
        const updatedData = allStorageData.filter((item, index) => index !== indexToDelete);
        setAllStorageData(updatedData);
    };

    const allFieldsFilled = () => {
        return (
            billInfo.billType &&
            billInfo.billname &&
            billInfo.month &&
            billInfo.year &&
            billInfo.amount
        );
    }
    const updateTypeName = (typeId) => {
        if (typeId === 1) {
            setTypeName("Elektrik Faturaları:");
        }
        else if (typeId === 2) {
            setTypeName("Su Faturaları:");
        }
        else if (typeId === 3) {
            setTypeName("Doğalgaz Faturaları:");
        }
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
                    setCarbonFootprint("")
                    setAnalyze("")
                    setAllStorageData([])
                    setTypeName("")
                    setResponseMessage('Faturanız Kaydedildi <br /> Karbon Ayakiziniz Güncellendi. <br /> Ortalama Fatura Tutarınız Güncellendi');
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    const showDB = (typeId) => {
        show.billType = typeId
        updateTypeName(typeId)
        fetch('http://localhost:3000/show', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(show),
        })
            .then(response => response.json())
            .then(data => {
                setCarbonFootprint("")
                setAnalyze("")
                setResponseMessage("")
                if (data !== -1) {
                    setAllStorageData(data);
                } else {
                    setAllStorageData([])
                }

            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    const carbon = () => {
        fetch('http://localhost:3000/carbon', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(cookies.customerData),
        })
            .then(response => response.json())
            .then(data => {
                setAllStorageData([])
                setTypeName("")
                setAnalyze("")
                setResponseMessage("")
                if (data !== "-1") {
                    setCarbonFootprint(data);
                } else {
                    setCarbonFootprint("")
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    const analiz = () => {
        fetch('http://localhost:3000/analiz', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(cookies.customerData),
        })
            .then(response => response.json())
            .then(data => {
                setAllStorageData([])
                setTypeName("")
                setCarbonFootprint("")
                setResponseMessage("")
                if (data !== "-1") {
                    setAnalyze(data);
                } else {
                    setAnalyze("")
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    const removeBill = (index) => {
        fetch('http://localhost:3000/remove', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(allStorageData[index]),
        })
            .then(response => response.json())
            .then(data => {
                if (data === 1) {
                    deleteItem(index)
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.')
            });
    }

    return (
        <div>
            <Header />
            <div className="info-page">
                <div className="left-panel">
                    <h2>Yeni Fatura Ekle</h2>
                    <div style={{ textAlign: "center" }}>

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
                                <option value="2">Su</option>
                                <option value="3">Doğalgaz</option>
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
                                placeholder="0₺"
                                step="10"
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
                        <button className="analyze-button" onClick={() => carbon()}>
                            Karbon Ayak İzini Hesapla
                        </button>
                        <button className="analyze-button" onClick={() => analiz()}>
                            Analiz Yap
                        </button>
                    </div>
                </div>
            </div>
            <div className="output-panel">
                <h3>Çıktı Ekranı</h3>
                <div className="output-content">
                    <div className="cart-header">
                        <h2>{typeName}</h2>
                    </div>
                    {allStorageData.map((item, index) => (
                        <div className="cart-item" key={index}>
                            <div className="cart-item-info">
                                <div className="basket-info">
                                    <div className="basket-details">
                                        <div className="basket-name">
                                            Fatura Adı: {item.billName}
                                        </div>
                                        <div className="basket-size">
                                            Yıl: {item.year} - Ay: {item.month}
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="cart-item-actions">
                                <div className="price">
                                    {item.amount}₺
                                </div>
                                <button
                                    className="delete-btn"
                                    onClick={() => removeBill(index)}
                                >
                                    Sil
                                </button>
                            </div>
                        </div>
                    ))}

                    {carbonFootprint !== "" && (
                        <>
                            <h2 style={{ color: 'red' }}>Karbon Ayak İzi</h2>
                            <h3>{carbonFootprint} kg CO₂</h3>
                        </>
                    )}
                    {Analyze !== "" && (
                        <>
                            <h2 style={{ color: 'red' }}>Analiz</h2>
                            <p
                                dangerouslySetInnerHTML={{ __html: Analyze }}
                            ></p>
                        </>
                    )}
                    {responseMessage !== "" && (
                        <>
                            <p
                                dangerouslySetInnerHTML={{ __html: responseMessage }}
                            ></p>
                        </>
                    )}
                </div>
            </div>
        </div>
    )
}

export default InfoPage