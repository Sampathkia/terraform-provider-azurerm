package odata

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShortType = string

const (
	ShortTypeAccessPackage                               ShortType = "accessPackage"
	ShortTypeAccessPackageAssignmentPolicy               ShortType = "accessPackageAssignmentPolicy"
	ShortTypeAccessPackageCatalog                        ShortType = "accessPackageCatalog"
	ShortTypeAccessPackageResourceRequest                ShortType = "accessPackageResourceRequest"
	ShortTypeAccessPackageQuestion                       ShortType = "accessPackageQuestion"
	ShortTypeAccessPackageTextInputQuestion              ShortType = "accessPackageTextInputQuestion"
	ShortTypeAccessPackageMultipleChoiceQuestion         ShortType = "accessPackageMultipleChoiceQuestion"
	ShortTypeAdministrativeUnit                          ShortType = "administrativeUnit"
	ShortTypeApplication                                 ShortType = "application"
	ShortTypeConditionalAccessPolicy                     ShortType = "conditionalAccessPolicy"
	ShortTypeConnectedOrganizationMembers                ShortType = "connectedOrganizationMembers"
	ShortTypeConnectionInfo                              ShortType = "connectionInfo"
	ShortTypeCountryNamedLocation                        ShortType = "countryNamedLocation"
	ShortTypeDevice                                      ShortType = "device"
	ShortTypeDirectoryRole                               ShortType = "directoryRole"
	ShortTypeDirectoryRoleTemplate                       ShortType = "directoryRoleTemplate"
	ShortTypeDomain                                      ShortType = "domain"
	ShortTypeEmailAuthenticationMethod                   ShortType = "emailAuthenticationMethod"
	ShortTypeExternalSponsors                            ShortType = "externalSponsors"
	ShortTypeFido2AuthenticationMethod                   ShortType = "fido2AuthenticationMethod"
	ShortTypeGroup                                       ShortType = "group"
	ShortTypeGroupMembers                                ShortType = "groupMembers"
	ShortTypeIpNamedLocation                             ShortType = "ipNamedLocation"
	ShortTypeInternalSponsors                            ShortType = "internalSponsors"
	ShortTypeNamedLocation                               ShortType = "namedLocation"
	ShortTypeMicrosoftAuthenticatorAuthenticationMethod  ShortType = "microsoftAuthenticatorAuthenticationMethod"
	ShortTypeOrganization                                ShortType = "organization"
	ShortTypePasswordAuthenticationMethod                ShortType = "passwordAuthenticationMethod"
	ShortTypePhoneAuthenticationMethod                   ShortType = "phoneAuthenticationMethod"
	ShortTypeRequestorManager                            ShortType = "requestorManager"
	ShortTypeServicePrincipal                            ShortType = "servicePrincipal"
	ShortTypeSingleUser                                  ShortType = "singleUser"
	ShortTypeSocialIdentityProvider                      ShortType = "socialIdentityProvider"
	ShortTypeTemporaryAccessPassAuthenticationMethod     ShortType = "temporaryAccessPassAuthenticationMethod"
	ShortTypeUser                                        ShortType = "user"
	ShortTypeWindowsHelloForBusinessAuthenticationMethod ShortType = "windowsHelloForBusinessAuthenticationMethod"
)

type Type = string

const (
	TypeAccessPackage                               Type = "#microsoft.graph.accessPackage"
	TypeAccessPackageAssignmentPolicy               Type = "#microsoft.graph.accessPackageAssignmentPolicy"
	TypeAccessPackageQuestion                       Type = "#microsoft.graph.accessPackageQuestion"
	TypeAccessPackageTextInputQuestion              Type = "#microsoft.graph.accessPackageTextInputQuestion"
	TypeAccessPackageMultipleChoiceQuestion         Type = "#microsoft.graph.accessPackageMultipleChoiceQuestion"
	TypeAccessPackageCatalog                        Type = "#microsoft.graph.accessPackageCatalog"
	TypeAccessPackageResourceRequest                Type = "#microsoft.graph.accessPackageResourceRequest"
	TypeAdministrativeUnit                          Type = "#microsoft.graph.administrativeUnit"
	TypeApplication                                 Type = "#microsoft.graph.application"
	TypeConditionalAccessPolicy                     Type = "#microsoft.graph.conditionalAccessPolicy"
	TypeConnectedOrganizationMembers                Type = "#microsoft.graph.connectedOrganizationMembers"
	TypeConnectionInfo                              Type = "#microsoft.graph.connectionInfo"
	TypeCountryNamedLocation                        Type = "#microsoft.graph.countryNamedLocation"
	TypeDevice                                      Type = "#microsoft.graph.device"
	TypeDirectoryRole                               Type = "#microsoft.graph.directoryRole"
	TypeDirectoryRoleTemplate                       Type = "#microsoft.graph.directoryRoleTemplate"
	TypeDomain                                      Type = "#microsoft.graph.domain"
	TypeEmailAuthenticationMethod                   Type = "#microsoft.graph.emailAuthenticationMethod"
	TypeExternalSponsors                            Type = "#microsoft.graph.externalSponsors"
	TypeFido2AuthenticationMethod                   Type = "#microsoft.graph.fido2AuthenticationMethod"
	TypeGroup                                       Type = "#microsoft.graph.group"
	TypeGroupMembers                                Type = "#microsoft.graph.groupMembers"
	TypeIpNamedLocation                             Type = "#microsoft.graph.ipNamedLocation"
	TypeInternalSponsors                            Type = "#microsoft.graph.internalSponsors"
	TypeNamedLocation                               Type = "#microsoft.graph.namedLocation"
	TypeMicrosoftAuthenticatorAuthenticationMethod  Type = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
	TypeOrganization                                Type = "#microsoft.graph.organization"
	TypePasswordAuthenticationMethod                Type = "#microsoft.graph.passwordAuthenticationMethod"
	TypePhoneAuthenticationMethod                   Type = "#microsoft.graph.phoneAuthenticationMethod"
	TypeRequestorManager                            Type = "#microsoft.graph.requestorManager"
	TypeServicePrincipal                            Type = "#microsoft.graph.servicePrincipal"
	TypeSingleUser                                  Type = "#microsoft.graph.singleUser"
	TypeSocialIdentityProvider                      Type = "#microsoft.graph.socialIdentityProvider"
	TypeTemporaryAccessPassAuthenticationMethod     Type = "#microsoft.graph.temporaryAccessPassAuthenticationMethod"
	TypeUser                                        Type = "#microsoft.graph.user"
	TypeWindowsHelloForBusinessAuthenticationMethod Type = "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod"
)
