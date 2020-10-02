import { CloseCircleOutlined, RightOutlined } from '@ant-design/icons';
import { Button, Card, Result } from 'antd';
import { FormattedMessage, formatMessage } from 'umi';
import React, { Fragment } from 'react';

import { GridContent } from '@ant-design/pro-layout';
import styles from './index.less';

const Content = (
  <Fragment>
    <div className={styles.title}>
      <FormattedMessage
        id="demotwo.error.hint-title"
        defaultMessage="The content you submitted has the following error:"
      />
    </div>
    <div style={{ marginBottom: 16 }}>
      <CloseCircleOutlined style={{ marginRight: 8 }} className={styles.error_icon} />
      <FormattedMessage
        id="demotwo.error.hint-text1"
        defaultMessage="Your account has been frozen"
      />
      <a style={{ marginLeft: 16 }}>
        <FormattedMessage id="demotwo.error.hint-btn1" defaultMessage="Thaw immediately" />
        <RightOutlined />
      </a>
    </div>
    <div>
      <CloseCircleOutlined style={{ marginRight: 8 }} className={styles.error_icon} />
      <FormattedMessage
        id="demotwo.error.hint-text2"
        defaultMessage="Your account is not yet eligible to apply"
      />
      <a style={{ marginLeft: 16 }}>
        <FormattedMessage id="demotwo.error.hint-btn2" defaultMessage="Upgrade immediately" />
        <RightOutlined />
      </a>
    </div>
  </Fragment>
);

export default () => (
  <GridContent>
    <Card bordered={false}>
      <Result
        status="error"
        title={formatMessage({ id: 'demotwo.error.title' })}
        subTitle={formatMessage({ id: 'demotwo.error.description' })}
        extra={
          <Button type="primary">
            <FormattedMessage id="demotwo.error.btn-text" defaultMessage="Return to modify" />
          </Button>
        }
        style={{ marginTop: 48, marginBottom: 16 }}
      >
        {Content}
      </Result>
    </Card>
  </GridContent>
);
