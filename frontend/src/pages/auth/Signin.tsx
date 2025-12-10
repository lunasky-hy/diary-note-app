import { Button, Card, Flex, Form, Input, Typography } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { useCallback, useState } from 'react';
const { Item: FormItem } = Form;
const { Title } = Typography;
const { Password } = Input;

type SigninForm = {
  username: string;
  password: string;
};

export default function Signin() {
  const [form] = useForm<SigninForm>();
  const [isSending, setIsSending] = useState(false);

  const handleSubmit = useCallback((values: SigninForm) => {
    setIsSending(true);
    fetch('/v1/api/auth/signin', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: values.username,
        password: values.password,
      }),
    })
      .then((raw) => raw.json())
      .then(() => setIsSending(false))
      .catch((e) => console.log(e));
  }, []);

  return (
    <Flex vertical gap={'middle'}>
      <Card>
        <Title level={2}>ユーザーログイン</Title>
        <Form form={form} layout="vertical" onFinish={handleSubmit}>
          <FormItem
            name={'username'}
            label="ユーザー名"
            rules={[{ required: true, message: '入力してください' }]}
          >
            <Input type={'text'} maxLength={20} />
          </FormItem>
          <FormItem
            name={'password'}
            label="パスワード"
            rules={[{ required: true, message: '入力してください' }]}
          >
            <Password type={'password'} />
          </FormItem>
          <FormItem>
            <Button type="primary" htmlType="submit" loading={isSending}>
              ログイン
            </Button>
          </FormItem>
        </Form>
      </Card>
    </Flex>
  );
}
