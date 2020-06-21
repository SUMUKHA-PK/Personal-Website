import React, { Component } from 'react';
import Contacts from './components/contacts';

class App extends Component {
  state = {
    contacts: []
  }

  componentDidMount() {
    fetch('http://127.0.0.1:44440/trips')
      .then(res => res.json())
      .then((data) => {
        this.setState({ contacts: data })
      })
      .catch(console.log)
  }

  render() {
    // return (
    //   <div class="card">
    //     <div class="card-body">
    //       <h5 class="card-title">Steve Jobs</h5>
    //       <h6 class="card-subtitle mb-2 text-muted">steve@apple.com</h6>
    //       <p class="card-text">Stay Hungry, Stay Foolish</p>
    //     </div>
    //   </div>
    // );
    return (
      <Contacts contacts={this.state.contacts} />
    );
  }
}

export default App;