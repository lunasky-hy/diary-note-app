import { Flex, Layout } from 'antd';
import { BrowserRouter, Route, Routes } from 'react-router';
import PostNote from './pages/PostNote';
import PostQuestions from './pages/PostQuestions';
import AppFooter from './layouts/Footer';

const layoutStyle: React.CSSProperties = {
  width: '100%',
  height: '100vh',
  textAlign: 'center',
};

const innerBodyStyle: React.CSSProperties = {
  width: '100%',
  maxWidth: 980,
  height: 'max-content',
  margin: 'auto',
};

function App() {
  return (
    <Layout hasSider={false} style={layoutStyle}>
      <Flex vertical style={innerBodyStyle}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<PostNote />} />
            <Route path="/questions" element={<PostQuestions />} />
          </Routes>
        </BrowserRouter>
      </Flex>
      <AppFooter />
    </Layout>
  );
}

export default App;
