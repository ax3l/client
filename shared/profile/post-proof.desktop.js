// @flow
import React from 'react'
import {Box, Button, CopyableText, Icon, LinkWithIcon, PlatformIcon, Text} from '../common-adapters'
import {globalStyles, globalColors, globalMargins} from '../styles'
import * as shared from './post-proof.shared'
import type {Props} from './post-proof'
import {clipboard} from 'electron'

const PostProof = (props: Props) => {
  const {
    platform, platformUserName, descriptionText, proofAction, onCancel,
    onCancelText, onComplete, isOnCompleteWaiting, errorMessage,
  } = props
  const {
    descriptionView, noteText, onCompleteText, proofText, platformSubtitle, proofActionIcon, proofActionText,
  } = shared.propsForPlatform(props)

  return (
    <Box style={styleContainer} onCopyCapture={(e) => {
      // disallow copying the whole screen by accident
      e.preventDefault()
      clipboard.writeText(proofText)
    }}>
      <Icon style={styleClose} type='iconfont-close' onClick={() => onCancel()} />
      {!!errorMessage && <Box style={styleErrorBanner}><Text style={styleErrorBannerText} type='BodySmallSemibold'>{errorMessage}</Text></Box>}
      <Box style={styleContentContainer}>
        <PlatformIcon platform={platform} overlay='icon-proof-unfinished' overlayColor={globalColors.grey} />
        <Text style={{...stylePlatformUsername, ...(stylePlatformSubtitle ? {} : {marginBottom: globalMargins.medium})}} type='Header'>{platformUserName}</Text>
        {!!platformSubtitle && <Text style={stylePlatformSubtitle} type='Body'>{platformSubtitle}</Text>}
        {descriptionView || (descriptionText && <Text type='Body'>{descriptionText}</Text>)}
        {!!proofText && <CopyableText style={styleProofText} value={proofText} />}
        {!!noteText && <Text style={styleNoteText} type='BodySmall'>{noteText}</Text>}
        {!!proofAction && !!proofActionText && !!proofActionIcon && <LinkWithIcon style={styleProofAction} label={proofActionText} icon={proofActionIcon} color={globalColors.blue} onClick={() => proofAction()} />}
        <Box style={styleButtonsContainer}>
          {!!onCancelText && <Button type='Secondary' onClick={() => onCancel()} label={onCancelText || 'Cancel'} />}
          <Button type='Primary' onClick={() => onComplete()} label={onCompleteText} waiting={isOnCompleteWaiting} />
        </Box>
      </Box>
    </Box>
  )
}

const styleContainer = {
  ...globalStyles.flexBoxColumn,
  flex: 1,
  alignItems: 'center',
  position: 'relative',
  paddingTop: globalMargins.large,
  paddingBottom: globalMargins.large,
  ...globalStyles.scrollable,
}

const styleClose = {
  position: 'absolute',
  top: globalMargins.small,
  right: globalMargins.small,
  ...globalStyles.clickable,
  color: globalColors.black_10,
}

const styleErrorBanner = {
  ...globalStyles.flexBoxColumn,
  justifyContent: 'center',
  alignItems: 'center',
  width: '100%',
  zIndex: 1,
  minHeight: globalMargins.large,
  padding: globalMargins.tiny,
  marginTop: -globalMargins.large,
  backgroundColor: globalColors.red,
}

const styleErrorBannerText = {
  color: globalColors.white,
  maxWidth: 512,
  textAlign: 'center',
}

const styleContentContainer = {
  ...globalStyles.flexBoxColumn,
  flex: 1,
  justifyContent: 'center',
  alignItems: 'center',
  margin: globalMargins.large,
  width: '100%',
  maxWidth: 512,
  textAlign: 'center',
}

const stylePlatformUsername = {
  color: globalColors.blue,
}

const stylePlatformSubtitle = {
  color: globalColors.black_10,
  marginBottom: globalMargins.medium,
}

const styleProofText = {
  width: '100%',
  minHeight: 116,
  maxHeight: 116,
  flexGrow: 1,
  marginTop: globalMargins.small,
}

const styleNoteText = {
  marginTop: globalMargins.tiny,
}

const styleProofAction = {
  marginTop: globalMargins.tiny,
}

const styleButtonsContainer = {
  ...globalStyles.flexBoxRow,
  marginTop: globalMargins.medium,
}

export default PostProof
