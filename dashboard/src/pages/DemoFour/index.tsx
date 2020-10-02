import { Button, Result } from 'antd';
import { FormattedMessage, formatMessage, Link } from 'umi';
import React from 'react';
import { RouteChildrenProps } from 'react-router';

import styles from './style.less';

const actions = (
  <div className={styles.actions}>
    <a href="">
      <Button size="large" type="primary">
        <FormattedMessage id="demofour.register-result.view-mailbox" />
      </Button>
    </a>
    <Link to="/">
      <Button size="large">
        <FormattedMessage id="demofour.register-result.back-home" />
      </Button>
    </Link>
  </div>
);

const DemoFour: React.FC<RouteChildrenProps> = ({ location }) => (
  <Result
    className={styles.registerResult}
    status="success"
    title={
      <div className={styles.title}>
        <FormattedMessage
          id="demofour.register-result.msg"
          values={{ email: location.state ? location.state.account : 'AntDesign@example.com' }}
        />
      </div>
    }
    subTitle={formatMessage({ id: 'demofour.register-result.activation-email' })}
    extra={actions}
  />
);

export default DemoFour;
