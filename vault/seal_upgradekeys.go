package vault

import (
	"context"
	"sync/atomic"
)

func (c *Core) SetUpgradeSealKeys() {
	atomic.StoreUint32(c.upgradeSealKeys, 1)
}

func (c *Core) doUpgradeSealKeys(ctx context.Context) {
	if err := c.seal.UpgradeKeys(ctx); err != nil {
		c.logger.Error("failed to upgrade seal keys", "error", err)
		return
	}
	atomic.StoreUint32(c.upgradeSealKeys, 0)
}
