import React, { Component } from 'react';
import  {Link} from 'react-router';
import sonataLogo from './sonata.png';
import './App.css';


class App extends Component {

  render() {
    return (
      <div className="container-fluid">
      <nav className="navbar navbar-toggleable-md navbar-inverse bg-inverse fixed-top">
        <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <a className="navbar-brand" href="/"> <img src="https://5g-ppp.eu/wp-content/uploads/2015/03/SONATA-new.png"
               alt="Sonata Logo"
               className="navbar-img"/></a>

        <div className="collapse navbar-collapse" id="sonataNavBar">
          <ul className="navbar-nav mr-auto">
            <li className="nav-item">
              <a className="nav-link"><Link to="/">Home</Link></a>
            </li>
            <li className="nav-item">
              <a className="nav-link"><Link to="/about">About</Link></a>
            </li>
          </ul>
        </div>
      </nav>


    <div className="parent">
      <div className="App-body">
        {this.props.children}
      </div>
    </div>


  </div>
    );
  }
}

export default App;
