import { Button, Card, Flex, Form, Input, Typography } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { useCallback, useState } from 'react';
const { Item: FormItem } = Form;
const { Title } = Typography;
const { Password } = Input;

type SignupForm = {
  username: string;
  password: string;
  confirm: string;
};

export default function Signup() {
  const [form] = useForm<SignupForm>();
  const [isSending, setIsSending] = useState(false);

  const handleSubmit = useCallback((values: SignupForm) => {
    setIsSending(true);
    fetch('/v1/api/auth/signup', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: values.username,
        password: values.password,
      }),
    })
      .then((raw) => raw.text())
      .then((d) => console.log(d))
      .then(() => setIsSending(false))
      .catch((e) => console.log(e));
  }, []);

  return (
    <Flex vertical gap={'middle'}>
      <Card>
        <Title level={2}>ユーザー登録</Title>
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
          <FormItem
            name={'confirm'}
            label="パスワード確認"
            rules={[
              { required: true, message: '入力してください' },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue('password') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('パスワードが一致しません'));
                },
              }),
            ]}
          >
            <Password type={'password'} />
          </FormItem>
          <FormItem>
            <Button type="primary" htmlType="submit" loading={isSending}>
              登録
            </Button>
          </FormItem>
        </Form>
      </Card>
    </Flex>
  );
}
