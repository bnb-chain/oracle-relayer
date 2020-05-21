pragma solidity 0.5.8;

contract TokenHubContract {
    event LogBindRequest(address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount);
    event LogBindSuccess(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount, uint256 decimals);
    event LogBindRejected(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol);
    event LogBindTimeout(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol);
    event LogBindInvalidParameter(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol);

    event LogTransferOut(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee);
    event LogBatchTransferOut(uint256 sequence, uint256[] amounts, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee);
    event LogBatchTransferOutAddrs(uint256 sequence, address[] recipientAddrs, address[] refundAddrs);

    event LogTransferInFailureTimeout(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime);
    event LogTransferInFailureInsufficientBalance(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 auctualBalance);
    event LogTransferInFailureUnboundToken(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol);
    event LogTransferInFailureUnknownReason(uint256 sequence, address refundAddr, address recipient, uint256 bep2TokenAmount, address contractAddr, bytes32 bep2TokenSymbol);
}