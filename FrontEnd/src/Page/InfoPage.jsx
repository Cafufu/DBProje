import React, { useEffect, useState } from 'react'
import { useCookies } from 'react-cookie';
import '../css/info.css'
import '../css/bill.css'
function InfoPage() {

    const [allStorageData, setAllStorageData] = useState({});
    const [cookies, setCookie, getCookies] = useCookies(['customerData']);

    const getAllLocalStorage = () => {
        let keys = Object.keys(localStorage);
        let localStorageData = {};

        keys.forEach(key => {
            localStorageData[key] = JSON.parse(localStorage.getItem(key));
        });
        return localStorageData;
    };

    const removeItemFromLocalStorage = (key) => {
        localStorage.removeItem(key);
        setAllStorageData(prevState => {
            const updatedState = { ...prevState };
            delete updatedState[key];
            return updatedState;
        });
    };

    const clearLocalStorage = () => {
        localStorage.clear();
        setAllStorageData({});
    };

    return (
        <div>
            <div className="info-page">

                <div className="left-panel">
                    <h2>Yeni Fatura Ekle</h2>
                    <form className="bill-form">

                        <div className="form-group">
                            <label htmlFor="bill-type">Fatura Türü</label>
                            <p>Çerez Değeri: {cookies.customerData || 'Henüz yok'}</p>
                            <select id="bill-type" name="bill-type">
                                <option value="electricity">Elektrik</option>
                                <option value="gas">Doğalgaz</option>
                                <option value="water">Su</option>
                            </select>
                        </div>


                        <div className="form-group">
                            <label htmlFor="bill-name">Fatura Adı</label>
                            <input type="text" id="bill-name" name="bill-name" placeholder="Faturanın adı" />
                        </div>


                        <div className="form-group row">
                            <div className="column">
                                <label htmlFor="month">Ay</label>
                                <select id="month" name="month">
                                    <option value="january">Ocak</option>
                                    <option value="february">Şubat</option>
                                    <option value="march">Mart</option>
                                    <option value="april">Nisan</option>
                                    <option value="may">Mayıs</option>
                                    <option value="june">Haziran</option>
                                    <option value="july">Temmuz</option>
                                    <option value="august">Ağustos</option>
                                    <option value="september">Eylül</option>
                                    <option value="october">Ekim</option>
                                    <option value="november">Kasım</option>
                                    <option value="december">Aralık</option>
                                </select>
                            </div>
                            <div className="column">
                                <label htmlFor="year">Yıl</label>
                                <input type="number" id="year" name="year" placeholder="2024" min="2000" max="2100" />
                            </div>
                        </div>


                        <div className="form-group">
                            <label htmlFor="amount">Fatura Tutarı</label>
                            <input type="number" id="amount" name="amount" placeholder="0.00" step="0.01" />
                        </div>


                        <button type="submit" className="btn-save">Kaydet</button>
                    </form>
                </div>

                <div className="right-panel">
                    <h2>Fatura Görüntüleme ve Analiz</h2>
                    <div className="button-group">
                        <button className="view-button">Elektrik Faturasını Görüntüle</button>
                        <button className="view-button">Su Faturasını Görüntüle</button>
                        <button className="view-button">Doğalgaz Faturasını Görüntüle</button>
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