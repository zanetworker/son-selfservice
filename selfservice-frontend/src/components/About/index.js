import React, {Component} from 'react'
import "./About.css"

class About extends Component {

render(){
  return(
    <div className="container">

      <h1 className="my-4">About Us

        <small>Welcome to the Self-Service Portal</small>
      </h1>
      <p></p>

      <div className="row">
        <div className="col-lg-12">
          <h2 className="my-4">Our Team</h2>
        </div>
        <div className="col-lg-4 col-sm-6 text-center mb-4">
          <img className="rounded-circle img-fluid d-block mx-auto" src="http://placehold.it/200x200" alt=""/>
          <h3>John Smith
            <small>Job Title</small>
          </h3>
          <p>What does this team member to? Keep it short! This is also a great spot for social links!</p>
        </div>
        <div className="col-lg-4 col-sm-6 text-center mb-4">
          <img className="rounded-circle img-fluid d-block mx-auto" src="http://placehold.it/200x200" alt=""/>
          <h3>John Smith
            <small>Job Title</small>
          </h3>
          <p>What does this team member to? Keep it short! This is also a great spot for social links!</p>
        </div>
        <div className="col-lg-4 col-sm-6 text-center mb-4">
          <img className="rounded-circle img-fluid d-block mx-auto" src="http://placehold.it/200x200" alt=""/>
          <h3>John Smith
            <small>Job Title</small>
          </h3>
          <p>What does this team member to? Keep it short! This is also a great spot for social links!</p>
        </div>
        <div className="col-lg-4 col-sm-6 text-center mb-4">
          <img className="rounded-circle img-fluid d-block mx-auto" src="http://placehold.it/200x200" alt=""/>
          <h3>John Smith
            <small>Job Title</small>
          </h3>
          <p>What does this team member to? Keep it short! This is also a great spot for social links!</p>
        </div>
      </div>

    </div>
  );
}
}


export default About
