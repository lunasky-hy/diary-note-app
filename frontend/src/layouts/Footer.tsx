import { Divider, Flex, Typography } from 'antd';
import { Footer } from 'antd/es/layout/layout';
const { Link } = Typography;

export default function AppFooter() {
  return (
    <Footer>
      <Flex gap={'middle'} justify="center">
        <Link href="/">日記を記録</Link>
        <Link href="/questions">質問を追加</Link>
      </Flex>
      <Divider />
      <Typography>Easy Dialy</Typography>
    </Footer>
  );
}
