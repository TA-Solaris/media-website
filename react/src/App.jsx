import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';

import Home from './pages/Home';

function App() {
  return (
    <section>
      <AppBar position="sticky">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Media Website
          </Typography>
        </Toolbar>
      </AppBar>
      <Home />
    </section>
  );
}

export default App;
